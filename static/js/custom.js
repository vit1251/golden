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

    summaryUpdateRoutine() {
        $.ajax({
            url: "/api/stat",
            type: "POST",
            dataType: "json",
            context: this,
            success: (resp) => {
                //
                console.log(resp);
                //
                const NetmailMessageCount = resp.NetmailMessageCount;
                const EchomailMessageCount = resp.EchomailMessageCount;
                const FileCount = resp.FileCount;
                //
                if (NetmailMessageCount > 0) {
                    $('#mainMenuDirect').show();
                    $('#mainMenuDirect').html(NetmailMessageCount);
                } else {
                    $('#mainMenuDirect').hide();
                }
                //
                if (EchomailMessageCount > 0) {
                    $('#mainMenuEcho').show();
                    $('#mainMenuEcho').html(EchomailMessageCount);
                } else {
                    $('#mainMenuEcho').hide();
                }
                //
                this.registerSummaryUpdateRoutine();
            }
        });
    }

    registerSummaryUpdateRoutine() {
        setTimeout(() => {
            this.summaryUpdateRoutine();
        }, 15000.0);
    }

    processNewPacketRoutine() {
        $.ajax({
            url: "/api/service/start",
            type: "POST",
            dataType: "json",
            context: this,
            data: {
                service: 'tosser',
            },
            success: (resp) => {
                console.log(resp);
                this.registerProcessNewPacketRoutine();
            }
        });
    }

    registerProcessNewPacketRoutine() {
        setTimeout(() => {
            this.processNewPacketRoutine();
        }, 30000.0);
    }

    run() {
        this.registerHandler();
        this.setupClock();
        this.registerKeyboard();
        this.summaryUpdateRoutine();
        this.processNewPacketRoutine();
    }

}

$(() => {
    const app = new Application();
    app.run();
});
