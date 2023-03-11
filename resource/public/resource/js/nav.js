
$(function () {
    $('#searchInput').keypress(function (event) {
        var keycode = (event.keyCode ? event.keyCode : event.which);
        if (keycode == '13') {
            let keyword = $(this).val()
            window.location.href = "/s/" + keyword + "/1";
        }
        event.stopPropagation();
    });
});