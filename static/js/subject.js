
function get_subject(){
    $.ajax({
        type: 'get',
        url: '/subject/list',
        success: function (res) {
            console.log(res)
            var info_list
            for(i=0;i<res.data.length;i++){
                info="<tr><td>"+res.data[i].id+"</td><td><u style=\"cursor:pointer\" class=\"text-primary\">"+res.data[i].content+"</u></td><td class=\"text-l\">"+res.data[i].answer+"</td><td><button>修改</button></td><td><button  onclick=\"delete_subject("+res.data[i].id+")\">删除</button></td> </tr>"
                info_list+=info
            }
            document.getElementById("table_data").innerHTML = info_list
        }
    })
}

window.onload = get_subject()


function delete_subject(id){
    let postData={
        id:parseInt(id)
    }
    $.ajax({
        type: "POST",
        url: "/subject/delete",
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
