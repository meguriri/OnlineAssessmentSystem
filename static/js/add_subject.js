function add_subject(){
    let postData={
        content:$('#content').val(),
        answer:$('#answer').val()
    }
    $.ajax({
        type: "POST",
        url: "/subject/create",
        data: JSON.stringify(postData),
        processData: false,
        contentType: false,
        dataType: "json",
        success: function (data) {
            console.log('create successfully!');
            if (data.code == 0) {
                window.location.href = '/subject';
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