
class ContentChangeFeature {

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

    register() {
    }

}

class MetricFeature {

    constructor(commandStream, cb) {
        this.commandStream = commandStream;
        this.monitoringInterval = 15 * 1000;
        this.cb = cb;
    }


    summaryUpdateRoutine() {
        console.log(`Update metrics...`);
        const request = {
            requestId: '',
            action: 'stat',
        };
        this.commandStream.send(request, (resp) => {
            if (this.cb) {
                this.cb(resp);
            }
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

class NotificationFeature {

    constructor(commandStream) {
        this.commandStream = commandStream;
        this.muted = true;
    }

    handleNotificationPermissionChange(permission) {
        console.log(`Notification permission is ${permission}`);
        if( permission === "granted" ) {
            this.muted = false;
        }
    }

    /**
     * Show message
     *
     */
    showMessage(body = 'Welcome in Golden Point application!') {
        const notify = new Notification("Golden Point", {
            tag: 'golden-point-event',
            body: body,
        });
    }

    register() {
        Notification.requestPermission(this.handleNotificationPermissionChange.bind(this));
    }
}

/**
 * Main entry point
 *
 */
const main = () => {

    const commandStream = new CommandStream();

    /* Register update content service */
    const contentChangeFeature = new ContentChangeFeature();
    contentChangeFeature.register();

    /* Register show notification service */
    const notificationFeature = new NotificationFeature(commandStream);
    notificationFeature.register();

    /* Register update counter service */
    const metricFeature = new MetricFeature(commandStream, (metrics) => {

        const {
            EchoMessageCount = 0,
            NetMessageCount = 0,
            FileCount = 0,
        } = metrics;

        /* Update content counter */
        contentChangeFeature.updateMetric('mainMenuDirect', NetMessageCount);
        contentChangeFeature.updateMetric('mainMenuEcho', EchoMessageCount);
        contentChangeFeature.updateMetric('mainMenuFile', FileCount);

        /* Show notification */
        if (!notificationFeature.mute) {
            if (NetMessageCount > 0) {
                notificationFeature.showMessage(`You have ${NetMessageCount} message(s)!`);
            }
        }

    });
    metricFeature.register();

};

main();

