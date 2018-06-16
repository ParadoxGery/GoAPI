var plot2 = null;

function fetchAjaxData(url, success) {
    $.ajax({
        url: url,
        dataType:"json",
        success: function(data) {
            success(data);
            console.log('loaded');
        }
    });
}

function createPlot(url) {
    fetchAjaxData(url, function(data) {
        if (plot2 == null) {
            plot2 = $.jqplot('chart2', data, {
                title: "AJAX JSON Data Renderer"
            });
        } else {
            plot2.replot({data: data});
            console.log('replotting');
        }
    });
}

$(document).ready(function(){
    var jsonurl = "./jsondata.txt";

    //Regenerate the plot on button click.
    $('#ajax-button').click(function() {
        createPlot(jsonurl);
    });

    //Generate the plot first time through.
    createPlot(jsonurl);
});