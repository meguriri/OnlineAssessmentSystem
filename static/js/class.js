var cid=0
function get_class(){
    $.ajax({
        type: 'get',
        url: '/class/list',
        success: function (res) {
            console.log(res)
            var info_list=""
            for(i=0;i<res.data.length;i++){
                info="<tr> <td>"+res.data[i].id+"</td><td><u style=\"cursor:pointer\" class=\"text-primary\">"+
                res.data[i].name+"</u></td><td class=\"text-l\">"+res.data[i].introduction+
                "</td><td>"+res.data[i].type+"</td><td><button  class=\"btn btn-success\" onclick=\"update_class("+res.data[i].id
                +")\">修改</button></td><td><button class=\"btn btn-danger\" onclick=\"delete_class("+res.data[i].id+")\">删除</button> </td></tr>"
                console.log(info)
                info_list+=info
            }
            document.getElementById("table_data").innerHTML = info_list
        }
    })
}

window.onload = get_class()


function delete_class(id){
    let postData={
        id:parseInt(id)
    }
    $.ajax({
        type: "POST",
        url: "/class/delete",
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

function update_class(id){
    console.log("id:",id)
    $('#updateModal').modal('show');
    cid=id
}

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

$(document).ready(function () {
    $('#upsubmit').click(function(){
        let postData={
            id:cid,
            name:$('#uname').val(),
            introduction:$('#uintroduction').val()
        }
        console.log("up post:",postData)
        $.ajax({
            type: "POST",
            url: "/class/update",
            data: JSON.stringify(postData),
            processData: false,
            contentType: false,
            dataType: "json",
            success: function (data) {
                console.log('update successfully!');
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
    })
})

