$(function() {

    $('.service-start').on('click', function(e) {
        //var serviceName = $''
        console.log('Start request');
        /* Start request */
        $.ajax({
            url: "/api/service/start",
            context: e,
            success: function() {
                $(this).html( "Всё ок" );
            }
        });
    });

});
