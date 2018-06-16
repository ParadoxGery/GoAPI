$(document).ready(function(){
    // Our ajax data renderer which here retrieves a text file.
    // it could contact any source and pull data, however.
    // The options argument isn't used in this renderer.
    var ajaxDataRenderer = function(url, plot, options) {
        var ret = null;
        $.get(url).done(function(data) {
            ret = data;
        })
        $.ajax({
            // have to use synchronous here, else the function
            // will return before the data is fetched
            async: false,
            crossDomain: true,
            url: url,
            dataType:"json",
            success: function(data) {
                ret = data;
            }
        });
        return ret;
    };

    // The url for our json data
    var jsonurl = "http://192.168.2.27:8888/temps";

    // passing in the url string as the jqPlot data argument is a handy
    // shortcut for our renderer.  You could also have used the
    // "dataRendererOptions" option to pass in the url.
    var plot2 = $.jqplot('chart1',jsonurl,{
        title: "Week",
        axes: {
            xaxis: {
                label: 'time',
                labelRenderer: $.jqplot.CanvasAxisLabelRenderer
            },
            yaxis: {
                label: 'temperature [°C]',
                labelRenderer: $.jqplot.CanvasAxisLabelRenderer
            }
        },
        dataRenderer: ajaxDataRenderer,
        dataRendererOptions: {
            unusedOptionalUrl: jsonurl
        }
    });
});