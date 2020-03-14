$(function() {

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

});
