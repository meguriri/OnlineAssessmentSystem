
subject_id_list=[20]
subject_number=0

function submit(){

    anwser_list=[subject_number]

    for(i=0;i<subject_number;i++){
           anwser_list[i]={
               id:subject_id_list[i],
               anwser:$('#id'+subject_id_list[i]).val()
        }
    }

    let post_data={
        anwser_list:anwser_list
    }

    $.ajax({
        type: 'post',
        url: '/test_paper/submit',
        dataType: 'json',
        data: JSON.stringify(post_data),
        success: function (res) {
            console.log(res)
            let info_list=""
            var success="#0bf116;\">√"
            var fail   ="#f1130b;\">×"
            for(i=0;i<res.data.gudge_anwser.length;i++){
                var result
                if(res.data.gudge_anwser[i].type!=0){
                    result=fail
                }else{
                    result=success
                }
                info="<tr><td>"+res.data.gudge_anwser[i].id+"</td><td><u style=\"cursor:pointer\" class=\"text-primary\">"+res.data.gudge_anwser[i].content+"</u></td><td class=\"text-l\">"+res.data.gudge_anwser[i].true_anwser+"</td> <td class=\"text-l\">"+res.data.gudge_anwser[i].user_anwser+"</td><td class=\"text-l\" style=\"color:"+result+"</td></tr>"
                info_list+=info
            }
            document.getElementById("table_data").innerHTML = info_list
        }
    })

}

function get_test_paper(){
    let postData={
        user_id:"admin",
        FacilityValue:"0"
    }

    $.ajax({
        type: 'post',
        url: '/test_paper/create',
        data: JSON.stringify(postData),
        processData: false,
        contentType: false,
        dataType: "json",
        success: function (res) {
            console.log(res)
            let info_list=""
            for(i=0;i<res.data.subject_list.length;i++){
                subject_number++
                subject_id_list[i]=res.data.subject_list[i].id
                info=" <tr> <td>"+res.data.subject_list[i].id+"</td><td><u style=\"cursor:pointer\" class=\"text-primary\">"+res.data.subject_list[i].content+"</u></td><td><input id=\"id"+res.data.subject_list[i].id+"\"></td></tr>"
                info_list+=info
            }
            document.getElementById("table_data").innerHTML = info_list
        }
    })

}

window.onload = get_test_paper()

