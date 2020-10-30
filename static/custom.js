
class MetricFeature {

    constructor() {
        this.summaryMonitoringInterval = 2.5 * 60 * 1000;
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

        fetch("/api/stat", {
            method: "POST",
            headers: {},
            body: '',
        })
        .then((response) => {
            return response.json();
        })
        .then((data) => {
            //
            console.log(data);
            //
            const NetmailMessageCount = data.NetmailMessageCount;
            const EchomailMessageCount = data.EchomailMessageCount;
            //const FileCount = resp.FileCount;
            //
            this.updateMetric('mainMenuDirect', NetmailMessageCount);
            this.updateMetric('mainMenuEcho', EchomailMessageCount);
            //
            this.registerSummaryUpdateRoutine();
        });
    }

    registerSummaryUpdateRoutine() {
        setTimeout(() => {
            this.summaryUpdateRoutine();
        }, this.summaryMonitoringInterval);
    }

    register() {
        this.registerSummaryUpdateRoutine();
    }

}

class ClockFeature {

    constructor() {
        this.sep = false;
    }

    makeTime() {
        const now = new Date();
        const min = now.getMinutes();
        const hour = now.getHours();
        const minStr = `${min}`.padStart(2, '0');
        const hourStr = `${hour}`.padStart(2, '0');
        if (this.sep) {
            const result = `${hourStr}:${minStr}`;
            return result;
        } else {
            const result = `${hourStr} ${minStr}`;
            return result;
        }
    }

    updateClock() {
        const currentTime = this.makeTime();
        let clock = document.getElementById("clock");
        clock.innerHTML = currentTime;
        this.sep = !this.sep;
    }

    setupClock() {
        setInterval(() => {
            this.updateClock();
        }, 1000);
    }

    register() {
        this.setupClock();
    }

}

class Application {

    constructor() {
        this.features = [
            new MetricFeature(),
            new ClockFeature(),
        ];
    }

    run() {
        this.features.forEach((feature) => {
            feature.register();
        });
    }

}

const app = new Application();
app.run();

