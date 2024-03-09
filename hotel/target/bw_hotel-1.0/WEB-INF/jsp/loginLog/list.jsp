<%@ page language="java" import="java.util.*" pageEncoding="UTF-8" %>
<%@ taglib uri="http://java.sun.com/jsp/jstl/core" prefix="c" %>
<%@ taglib prefix="fmt" uri="http://java.sun.com/jsp/jstl/fmt" %>

<c:set value="${pageContext.request.contextPath}" scope="page" var="ctx"></c:set>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
<head>

    <title></title>

    <meta http-equiv="pragma" content="no-cache">
    <meta http-equiv="cache-control" content="no-cache">
    <meta http-equiv="expires" content="0">
    <meta http-equiv="keywords" content="keyword1,keyword2,keyword3">
    <meta http-equiv="description" content="This is my page">
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>


    <!--   <link rel="stylesheet" href="${ctx}/css/roomset/roomset.css" type="text/css"></link> -->
    <link rel="stylesheet" href="${ctx}/bootstrap/css/bootstrap.css" type="text/css"></link>
    <link rel="stylesheet" href="${ctx}/css/page.css" type="text/css"></link>
    <link href="${ctx}/bootstrap/css/bootstrap-responsive.css" rel="stylesheet">  <!-- start 响应式布局要添加的 -->
    <script src="${ctx}/bootstrap/js/jquery-3.1.1.min.js"></script>
    <script src="${ctx}/bootstrap/js/bootstrap.js"></script>
    <script type="text/javascript" src="${ctx}/js/page.js"></script>
    <style>

        .container {
            margin-top: 10px;
        }

        .labelroomnumber {
            position: relative;
            font-size: 17px;
            float: left;
            margin-top: 15px;
        }

        .textone {
            margin-top: 12px;
        }

        .rightOne {
            margin-right: 50px;
            font-size: 16px;
        }

        .table th, .table td {
            text-align: center;
        }

        .theadone {
            background-color: #438eb9;
            color: #fff;
        }

        .dgvone {
            margin-top: 12px;
        }

        .roomnumberwidth {
            width: 70%;
        }


    </style>

</head>


<body>
<div class="container">
    <div class="span5">
        <div class="row-fluid">
            <label class="labelroomnumber">登录名称：</label>
            <form action="${ctx}/loginLog/tolist.do" method="post" style="float: left;">
                <input id="txtnameid" name="txtname" class="textone roomnumberwidth"
                       style="border-radius:0px; border-top-left-radius:4px; border-bottom-left-radius:4px;height:26px;"
                       type="text" placeholder="请输入关键字" value="${txtname}">
                <div class="input-append">
                    <button type="submit" class="btn-success textone" style="margin-left:-4px;height:26px;">
                        <li class="icon-search icon-white"></li>
                        搜索
                    </button>
                </div>
            </form>
        </div>
    </div>
    <div class="span6">
        <div class="row-fluid">
            <div class="span3">
                <button class="btn btn-info btn-small textone" type="button" onclick="addfunction()" disabled>
                    <li class="icon-plus icon-white"></li>
                    新增
                </button>
            </div>
            <div class="span3">
                <button class="btn btn-warning btn-small textone" type="button" onclick="updatefunction()" disabled>
                    <li class="icon-pencil icon-white"></li>
                    修改
                </button>
            </div>
            <div class="span3">
                <%-- <button class="btn btn-danger btn-small textone" type="button" onclick="deletefunction()">
                     <li class="icon-remove icon-white"></li>
                     删除  class="btn btn-primary btn-lg"
                 </button>--%>
                <!-- Button trigger modal -->
                <button type="button" class="btn btn-danger btn-small textone" data-toggle="modal"
                        data-target="#myModal">
                    <li class="icon-remove icon-white"></li>
                    删除
                </button>

                <!-- Modal -->
                <div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
                    <div class="modal-dialog" role="document">
                        <div class="modal-content">
                            <div class="modal-header">
                                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span
                                        aria-hidden="true">&times;</span></button>
                                <h4 class="modal-title" id="myModalLabel">超级密码</h4>
                            </div>
                            <div class="modal-body">
                                <input type="password" id="adminpwd" placeholder="请输入密码" name="superpwd" value=""/>
                            </div>
                            <div class="modal-footer">
                                <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
                                <button type="button" class="btn btn-primary" onclick="deletefunction()">提 交</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <br>
    <div class="dgvone">
        <table class="table table-condensed table-bordered table-striped" id="tableid">
            <thead class="theadone">
            <tr>
                <th>序号</th>
                <th>管理员名称</th>
                <th>登陆时间</th>
                <th>退出时间</th>
            </tr>
            </thead>
            <tbody id="tbody">
            <c:forEach items="${list.result}" var="item">
                <tr>
                    <td><input type="checkbox" name="id" value="${item.id}"></td>
                    <td>${item.loginname}</td>
                    <td>${item.logintime}</td>
                    <td>${item.exittime}</td>
                </tr>
            </c:forEach>
            </tbody>
        </table>
    </div>

    <div class="span11">
        <div class="row-fluid">
            <div class="tcdPageCode" style="text-align:center;"></div>
        </div>
    </div>
</div>


<script type="text/javascript">
    function addfunction() {
        <%--parent.document.getElementById('Mainid').src = '${ctx}/loginLog/toadd.do';--%>
    }

    function updatefunction() {
        var chk_value = [];
        $('input[name="id"]:checked').each(function () {
            chk_value.push($(this).val());
        });
        if (chk_value != "") {
            if (chk_value.toString().indexOf(",") > 0) {
                alert("修改只能选择一条");
            } else {
                var userName = parent.document.getElementById('loginName').innerText;
                $.ajax({
                    cache: false,                                             //是否使用缓存提交 如果为TRUE 会调用浏览器的缓存 而不会提交
                    type: "POST",                                           //上面3行都是必须要的
                    url: '${ctx}/loginLog/findLoginNameByuserName.do',              //地址 type 带参数
                    data: {"userName": userName, "userId": chk_value[0]},                    // IDCardValue 自定义的。相当于name把值赋予给 他可以在servlet 获取
                    async: false,                                          // 是否 异步 提交
                    success: function (result) {                          // 不出现异常 进行立面方
                        if (result == 1) {
                            parent.document.getElementById("Mainid").src = '${ctx}/loginLog/toupdate.do?id=' + chk_value[0];
                        } else {
                            $('input[name="id"]:checked').prop("checked", "");
                            alert("您不能需改他人的信息！！！");
                        }
                    },
                    error: function (data) {
                    }
                });


            }
        } else {
            alert("请选择一条数据进行修改");
        }
    }

    function deletefunction() {
        var pwd = $("#adminpwd").val();
        $.ajax({
            cache:false,
            type:"POST",
            url:'${ctx}/loginLog/isAdmin.do',
            data:{'adminPwd':pwd},
            async: false,
            success:function (result) {
                if (result){
                    var chk_value = [];
                    $('input[name="id"]:checked').each(function () {
                        chk_value.push($(this).val());
                    });
                    if (chk_value != "") {
                        var flag = window.confirm("注意：您确定要永久删除该信息吗?");
                        if (flag) {
                            parent.document.getElementById("Mainid").src = '${ctx}/loginLog/delete.do?ids=' + chk_value;
                        }
                    } else {
                        alert("请选择一条或多条数据进行删除");
                    }
                }else{
                    alert("您没有此权限");
                    $("#adminpwd").val("");
                }
            }
        });
    }

    /* 分页要用的 */
    $(".tcdPageCode").createPage({
        pageCount:${list.totalPage},
        current:${list.currentPage},
        backFn: function (p) {
            var txtname = document.getElementById("txtnameid").value;
            location.href = "${ctx}/loginLog/tolist.do?currentPage=" + p + "&txtname=" + txtname;
        }
    });

</script>

</body>
</html>
