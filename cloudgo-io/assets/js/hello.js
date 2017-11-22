$(document).ready(function() {
    $.ajax({
        url: "/title"
    }).then(function(data) {
       $('.title').append(data.title);
    });
});
