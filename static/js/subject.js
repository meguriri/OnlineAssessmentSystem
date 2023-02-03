var sid=0
function get_subject(){
    $.ajax({
        type: 'get',
        url: '/subject/list',
        success: function (res) {
            console.log(res)
            var info_list=""
            for(i=0;i<res.data.length;i++){
                info="<tr><td>"+res.data[i].id+"</td><td><u style=\"cursor:pointer\" class=\"text-primary\">"+
                res.data[i].content+"</u></td><td class=\"text-l\">"+res.data[i].answer+
                "</td><td><button onclick=\"update_subject("+res.data[i].id
                +")\"  class=\"btn btn-success\">修改</button></td><td><button  class=\"btn btn-danger\" onclick=\"delete_subject("+res.data[i].id+")\">删除</button></td> </tr>"
                info_list+=info
                console.log(info)
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

function update_subject(id){
    console.log("id:",id)
    $('#updateModal').modal('show');
    sid=id
}

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

$(document).ready(function () {
    $('#upsubmit').click(function(){
        let postData={
            id:sid,
            content:$('#ucontent').val(),
            answer:$('#uanswer').val()
        }
        console.log("up post:",postData)
        $.ajax({
            type: "POST",
            url: "/subject/update",
            data: JSON.stringify(postData),
            processData: false,
            contentType: false,
            dataType: "json",
            success: function (data) {
                console.log('update successfully!');
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
    })
})