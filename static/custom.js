
class MetricFeature {

    constructor(commandStream) {
        this.commandStream = commandStream;
        this.monitoringInterval = 15 * 1000;
    }

    updateMetric(metricName, value) {
        let element = document.getElementById(metricName);
        //
        element.innerHTML = value;
        //
        if (value > 0) {
            element.classList.remove("hidden");
        } else {
            element.classList.add("hidden");
        }
    }

    summaryUpdateRoutine() {
        console.log(`Update metrics...`);
        const request = {
            requestId: '',
            action: 'stat',
        };
        this.commandStream.send(request, (resp) => {
            /* Process response */
            const {
                EchoMessageCount = 0,
                NetMessageCount = 0,
                FileCount = 0,
            } = resp;
            /* Update parameters */
            this.updateMetric('mainMenuDirect', NetMessageCount);
            this.updateMetric('mainMenuEcho', EchoMessageCount);
            this.updateMetric('mainMenuFile', FileCount);
        });
        /* Register next update */
        this.registerSummaryUpdateRoutine();
    }

    registerSummaryUpdateRoutine() {
        setTimeout(() => {
            this.summaryUpdateRoutine();
        }, this.monitoringInterval);
    }

    register() {
        this.registerSummaryUpdateRoutine();
    }

}

class CommandStream {

    constructor() {
        const url = new URL('/api/v1', window.location.href);
        url.protocol = 'ws';
        this.socket = new WebSocket(url.href);
        this.socket.onopen = () => {
            this.state = 'READY';
        };
        this.socket.onmessage = (event) => {
            const {data = ''} = event;
            const msg = JSON.parse(data);
            console.log(`Incoming message: ${JSON.stringify(msg)}`);
            if (this.resolve) {
                this.resolve(msg);
                this.resolve = null;
            }
        };
        this.socket.onclose = () => {
            console.log(`Session is close.`);
            this.state = 'CLOSE';
        };
    }

    send(options, resolve) {
        this.resolve = resolve;
        this.socket.send(JSON.stringify(options));
    }

}

/**
 * Main entry point
 *
 */
const main = () => {

    const commandStream = new CommandStream();

    /* Register metric update */
    const metricFeature = new MetricFeature(commandStream);
    metricFeature.register();

};

main();

