
function get_class(){
    $.ajax({
        type: 'get',
        url: '/class/list',
        success: function (res) {
            console.log(res)
            var info_list
            for(i=0;i<res.data.length;i++){
                info="<tr> <td>"+res.data[i].id+"</td><td><u style=\"cursor:pointer\" class=\"text-primary\">"+res.data[i].name+"</u></td><td class=\"text-l\">"+res.data[i].introduction+"</td><td>"+res.data[i].type+"</td><td><button onclick=\"update_class("+res.data[i].id+")\">修改</button></td><td><button onclick=\"delete_class("+res.data[i].id+")\">删除</button> </td></tr>"
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
    prompt("课程名","介绍","添加授课教师")
}