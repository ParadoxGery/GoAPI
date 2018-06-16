var plot1 = null;
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

function createPlot(url, place, headline, yAxisLabel) {
    fetchAjaxData(url, function(data) {
        plot = $.jqplot(place, data, {
            title: headline,
            axes: {
                xaxis: {
                    label: 'time [hh:mm]',
                    renderer:$.jqplot.DateAxisRenderer,
                    //tickOptions:{formatString:'%H:%M'},
                    //tickInterval:'1 hour'
                },
                yaxis: {
                    label: yAxisLabel,
                    labelRenderer: $.jqplot.CanvasAxisLabelRenderer
                }
            }
        });
    });
}

$(document).ready(function(){
    createPlot("http://localhost:8888/tempdata","chart1","Day Temperature","temperature [°C]");
    createPlot("http://localhost:8888/hudata","chart2","Day Humidity","humidity [%]");
});