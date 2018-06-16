$(document).ready(function(){
    $("#slidy").click(function(){
        var Enabled = $("#check").prop('checked');
        //console.log(Enabled)
        if (Enabled) {
            $.ajax({
                url: '/io/10',
                type: 'DELETE',
            });
        }
        else {
            $.ajax({
                url: '/io/10',
                type: 'PUT',   
            });
        }
    
    });
});