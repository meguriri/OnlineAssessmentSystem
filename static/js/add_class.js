function add_class(){
    let postData={
        name:$('#name').val(),
        introduction:$('#introduction').val()
    }
    $.ajax({
        type: "POST",
        url: "/class/create",
        data: JSON.stringify(postData),
        processData: false,
        contentType: false,
        dataType: "json",
        success: function (data) {
            console.log('create successfully!');
            if (data.code == 0) {
                window.location.href = '/class';
            } else {
                alert(data.msg)
                console.log('login failed!');
            }
        },
        error: function () {
            alert('submit failed!');
        },
    })
}