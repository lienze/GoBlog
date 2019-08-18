
function getquerystring(name){
    var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    if(r!=null)return  unescape(r[2]); return null;
}
function clickcomment(){
    var name = document.getElementById("FormName");
    var comment = document.getElementById("FormComment");
    var postname = getquerystring("name");
    //alert(postname);
    var pageAddr = "upcomment?name="+name.value+"&comment="+comment.value+"&postname="+postname;
    window.location.href = pageAddr;
}
function initpage(){
    convert2md();
}
function convert2md(){
    var list = document.getElementsByName("markdownContent");
    for(var i=0; i < list.length; i++){
        compile(list[i]);
    }
}
function compile(elem){
    var converter = new showdown.Converter();
    var html = converter.makeHtml(elem.textContent);
    elem.innerHTML = html;
}