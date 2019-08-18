
function initpage(){
    alert("upload comment succeed!");
    var postname = getquerystring("postname");
    var pageAddr = "showpost?name=" + postname;
    window.location.href = pageAddr;
}
function getquerystring(name){
    var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    if(r!=null)return  unescape(r[2]); return null;
}