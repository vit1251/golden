class Application {

    constructor() {
        this.chart = null;
    }

    updateChart() {
        $.ajax({
            url: '/stat/image',
            method: 'GET',
            success: (resp) => {
                const labels = [];
                const serie1 = [];
                /* Parse response */
                resp.forEach((item) => {
                    const date = item.Date;
                    const value = item.Value;
                    labels.push(date);
                    serie1.push(value);
                });
                /* Update chart */
                console.log(resp);
                this.chart.update({
                    labels: labels,
                    series: [
                        serie1
                    ]
                });
            }
        });

    }

    renderChart(points) {
        const data = {
            labels: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri'],
            series: [
                [0, 0, 0, 0, 0]
            ]
        };
        const options = {
            width: 640,
            height: 480,
            showArea: true,
            showLine: false,
            showPoint: true,
        };
        this.chart = new Chartist.Line('.ct-chart', data, options);
        //new Chartist.Bar('.ct-chart', {
        //    labels: ['Понедельник', 'Вторник', 'Среда', 'Четверг', 'Пятница', 'Суббота', 'Воскресенье'],
        //    series: [
        //        [1, 2, 3, 4, 5, 6, 7],
        //    ]
        //}, options);
    }

    registerHandler() {
        $('.service-start').on('click', (e) => {
        let currentTarget = e.currentTarget;
        let serviceName = currentTarget.dataset.service;
        console.log('Start request: servcie = ', serviceName);
        /* Start request */
        $.ajax({
            url: "/api/service/start",
            type: "POST",
            dataType: "json",
            context: currentTarget,
            data: {
                service: serviceName,
            },
            success: function() {
                $(this).html( "Всё ок" );
            }
        });
        });
    }

    updateClock() {
        let now = new Date();
        let clock = document.getElementById("clock");
        clock.innerHTML = now.toLocaleTimeString();
    }

    setupClock() {
        setInterval(() => {
            this.updateClock();
        }, 1000);
    }

    run() {
        this.registerHandler();
        this.setupClock();
        this.renderChart();
        this.updateChart();
    }

}

$(() => {
    const app = new Application();
    app.run();
});
