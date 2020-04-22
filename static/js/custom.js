class Application {

    constructor() {
        this.chart = null;
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

    registerKeyboard() {
        Mousetrap.bind('ctrl+left', () => {
            console.log('Search prev message');
        });
        Mousetrap.bind('ctrl+right', () => {
            console.log('Search next message');
        });
    }

    run() {
        this.registerHandler();
        this.setupClock();
        this.registerKeyboard();
    }

}

$(() => {
    const app = new Application();
    app.run();
});
