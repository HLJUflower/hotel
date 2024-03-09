<%@ page language="java" import="java.util.*" pageEncoding="UTF-8" %>
<%@ taglib uri="http://java.sun.com/jsp/jstl/core" prefix="c" %>
<%@ taglib prefix="fmt" uri="http://java.sun.com/jsp/jstl/fmt" %>
<c:set value="${pageContext.request.contextPath}" scope="page" var="ctx"></c:set>
<%
    String path = request.getContextPath();
    String basePath = request.getScheme() + "://" + request.getServerName() + ":" + request.getServerPort() + path + "/";
%>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
<head>
    <title>首页</title>
    <meta http-equiv="pragma" content="no-cache">
    <meta http-equiv="cache-control" content="no-cache">
    <meta http-equiv="expires" content="0">
    <meta http-equiv="keywords" content="keyword1,keyword2,keyword3">
    <meta http-equiv="description" content="This is my page">
    <!--
    <link rel="stylesheet" type="text/css" href="styles.css">
    -->
    <!--   <link rel="stylesheet" href="${ctx}/css/roomset/roomset.css" type="text/css"></link> -->
    <link rel="stylesheet" href="${ctx}/bootstrap/css/bootstrap.css" type="text/css"></link>
    <link href="${ctx}/bootstrap/css/bootstrap-responsive.css" rel="stylesheet">  <!-- start 响应式布局要添加的 -->
    <script src="${ctx}/bootstrap/js/jquery-3.1.1.min.js"></script>
    <script src="${ctx}/bootstrap/js/bootstrap.js"></script>
    <link rel="stylesheet" href="${ctx}/bootstrap/css/bootstrap.min.css" type="text/css"></link>

    <script type="text/javascript" src="${ctx}/My97DatePicker/WdatePicker.js"></script>
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

        .marginrightone {
            margin-right: 33px;
        }

        .marginrighttwo {
            margin-right: 16.5px;
        }

        .marginrightthree {
            margin-right: 43px;
        }

        .textone {
            margin-top: 12px;
        }

        .inputone {
            width: 54.2%;
        }

        .inputtwo {
            width: 46.8%;
        }

        .inputthree {
            width: 46.8%;
        }

        .cboone {
            margin-top: 10px;
            height: 27px;
        }

        .margin-top-one {
            margin-top: -10px;
        }

        .margin-top-two {
            margin-top: -20px;
        }

        .textwidth {
            width: 138px;
        }

        .radiusone {
            margin-left: -4px;
        }

        @media (min-width: 731px) and (max-width: 767px) {
            .inputthree {
                width: 50.5%;
            }
        }

        @media (min-width: 550px) and (max-width: 730px) {
            .inputtwo {
                width: 46.2%;
            }

            .inputthree {
                width: 49.2%;
            }
        }


        @media (min-width: 431px) and (max-width: 550px) {
            .inputtwo {
                width: 43.2%;
            }

            .inputthree {
                width: 47.5%;
            }
        }

        @media (min-width: 366px) and (max-width: 430px) {
            .inputtwo {
                width: 40.2%;
            }

            .inputthree {
                width: 46%;
            }
        }

        @media (min-width: 285px) and (max-width: 366px) {
            .inputtwo {
                width: 37.2%;
            }

            .inputthree {
                width: 44%;
            }
        }

        @media (min-width: 237px) and (max-width: 285px) {
            .inputtwo {
                width: 30%;
            }

            .inputthree {
                width: 40%;
            }
        }

        .span12 {
            width: 95%;

        }

        .yansered {
            color: red;
        }

    </style>
</head>
<body style="background-color: #bbc1c6">
<div class="container" style="height:1000px;pxoverflow-x:auto;background-color: rgba(203,207,121,0.45)">

    <div class="span11" style=" border: solid; border-color: #DDDDDD;">
        <div class="span9 margin-top-one">
            <div class="row-fluid">
                <h3 style="text-align: center;color: rgba(0,54,190,0.79);margin-top: 15px">酒店客房管理系统</h3>
                <h3 style="text-align: right;">
                    <button style="background-color:#c8942d" onclick="login()">登录</button>
                </h3>
            </div>
        </div>

        <%-----------------------------------%>
        <form id="form1" method="post">
            <input type="hidden" name="receiveTargetID" value="${tuanDuiID}"/>
            <input id="roomId" name="roomID" type="hidden">
            <input id="LvKeLeiXingId" type="hidden" value="${LvKeLeiXingId}">
            <!--  ———————————————————————————————————————————————————————————————————————————————————————— -->
            <div class="span12">
                <div class="row-fluid">
                    <div class="span3">
                        <label style="font-weight: bolder;font-size: 15px">房间具体描述：</label>
                    </div>
                </div>
            </div>
            <div class="span12">
                <div class="row-fluid">
                    <div class="span3">
                        <label>房间号：</label>
                        <input id="roomNameId" name="roomName" type="text" style="width:100%;height:27px;"
                               readonly="readonly">
                    </div>
                    <div class="span3">
                        <label>标准房价：</label>
                        <input id="biaoZhunFangJiaId" type="text" style="width:100%;height:27px;" readonly="readonly">
                    </div>
                    <div class="span3">
                        <label>房间床位：</label>
                        <input id="countBed" type="text" style="width:100%;height:27px;" readonly="readonly">
                    </div>
                </div>
            </div>
            <!--  ———————————————————————————————————————————————————————————————————————————————————————— -->
            <div class="span12">
                <div class="row-fluid">

                    <div class="span3">
                        <label>钟点房间/小时：</label>
                        <input id="zhongDianFangId" type="text" style="width:100%;height:27px;" readonly="readonly">
                    </div>
                    <div class="span3">
                        <label>首段价格：</label>
                        <input id="shouDuanJiaGeId" type="text" style="width:100%;height:27px;" readonly="readonly">
                    </div>
                    <div class="span3">
                        <label>首段时长：</label>
                        <input id="shouDuanShiChangId" type="text" style="width:100%;height:27px;" readonly="readonly">
                    </div>
                </div>
            </div>
        </form>
        <!--  ———————————————————————————————————————————————————————————————————————————————————————— -->
        <div class="span12">
            <fieldset>
                <legend>房间状态：</legend>
            </fieldset>
        </div>

        <div class="span12" style="margin-top:-17px;">
            <div class="row-fluid">
                <button class="btn btn-info btn-small" type="button" onclick="allroomfunction()">所有房间</button>
                <button class="btn btn-info btn-small" type="button" onclick="onefunction()">大床房</button>
                <button class="btn btn-info btn-small" type="button" onclick="twofunction()">单人普通间</button>
                <button class="btn btn-info btn-small" type="button" onclick="threefunction()">二人普通房</button>
                <button class="btn btn-info btn-small" type="button" onclick="fourfunction()">二人标准间</button>
                <button class="btn btn-info btn-small" type="button" onclick="fivefunction()">豪华单间</button>
                <button class="btn btn-info btn-small" type="button" onclick="fiveTwofunction()">豪华双人间</button>
                <button class="btn btn-info btn-small" type="button" onclick="eightfunction()">浴缸房</button>
                <button class="btn btn-info btn-small" type="button" onclick="moviefunction()">影视房</button>
                <button class="btn btn-info btn-small" type="button" onclick="sixfunction()">会议室</button>
                <button class="btn btn-info btn-small" type="button" onclick="sevenfunction()">总统套房</button>
            </div>
        </div>

        <div class="span12" id="dynamicid" style="height:650px;overflow-x:auto;">
            <div class="row-fluid" id="div1">
                <c:forEach items="${list}" var="item">
                    <c:if test="${item.roomStateID==1}">
                        <button onclick="suibian(this)"
                                style="width:125px;height:110px;border: 3px solid #666666;  float:left;margin:2px; background:rgba(19,226,40,0.61);">
                            <input style="display: none;" value="${item.id}"/>
                            <input style="display: none;" value="${item.roomStateID}"/>
                            <label>${item.roomNumber}</label>
                            <label style="margin-top:-5px;">${item.roomName}</label>
                            <label style="margin-top:-5px;">${item.guestRoomLevelName}</label>
                            <label style="margin-top:-5px;">￥${item.standardPriceDay}</label>
                            <input style="display: none;" value="${item.standardPrice}"/>
                            <input style="display: none;" value="${item.firstPrice}"/>
                            <input style="display: none;" value="${item.firstDuration}"/>
                            <input style="display: none;" value="${item.roomAmount}"/>
                        </button>
                    </c:if>
                    <c:if test="${item.roomStateID==2}">
                        <button onclick="suibian(this)"
                                style="width:125px;height:110px;border: 3px solid #666666; float:left;margin:2px; background:#e3bfbf;">
                            <input style="display: none;" value="${item.id}"/>
                            <input style="display: none;" value="${item.roomStateID}"/>
                            <label>${item.roomNumber}</label>
                            <label style="margin-top:-5px;">${item.roomName}</label>
                            <label style="margin-top:-5px;">${item.guestRoomLevelName}</label>
                            <label style="margin-top:-5px;">￥${item.standardPriceDay}</label>
                            <input style="display: none;" value="${item.standardPrice}"/>
                            <input style="display: none;" value="${item.firstPrice}"/>
                            <input style="display: none;" value="${item.firstDuration}"/>
                            <input style="display: none;" value="${item.roomAmount}"/>
                        </button>
                    </c:if>
                    <c:if test="${item.roomStateID==4}">
                        <button onclick="suibian(this)"
                                style="width:125px;height:110px;border: 3px solid #666666; float:left;margin:2px; background:#63cdcd;">
                            <input style="display: none;" value="${item.id}"/>
                            <input style="display: none;" value="${item.roomStateID}"/>
                            <label>${item.roomNumber}</label>
                            <label style="margin-top:-5px;">${item.roomName}</label>
                            <label style="margin-top:-5px;">${item.guestRoomLevelName}</label>
                            <label style="margin-top:-5px;">￥${item.standardPriceDay}</label>
                            <input style="display: none;" value="${item.standardPrice}"/>
                            <input style="display: none;" value="${item.firstPrice}"/>
                            <input style="display: none;" value="${item.firstDuration}"/>
                            <input style="display: none;" value="${item.roomAmount}"/>
                        </button>
                    </c:if>
                    <c:if test="${item.roomStateID==5}">
                        <button onclick="suibian(this)"
                                style="width:125px;height:110px;border: 3px solid #666666; float:left;margin:2px; background:#d0d00e;">
                            <input style="display: none;" value="${item.id}"/>
                            <input style="display: none;" value="${item.roomStateID}"/>
                            <label>${item.roomNumber}</label>
                            <label style="margin-top:-5px;">${item.roomName}</label>
                            <label style="margin-top:-5px;">${item.guestRoomLevelName}</label>
                            <label style="margin-top:-5px;">￥${item.standardPriceDay}</label>
                            <input style="display: none;" value="${item.standardPrice}"/>
                            <input style="display: none;" value="${item.firstPrice}"/>
                            <input style="display: none;" value="${item.firstDuration}"/>
                            <input style="display: none;" value="${item.roomAmount}"/>
                        </button>
                    </c:if>
                    <c:if test="${item.roomStateID==6}">
                        <button onclick="suibian(this)"
                                style="width:125px;height:110px;border: 3px solid #666666; float:left;margin:2px; background:#e57044;">
                            <input style="display: none;" value="${item.id}"/>
                            <input style="display: none;" value="${item.roomStateID}"/>
                            <label>${item.roomNumber}</label>
                            <label style="margin-top:-5px;">${item.roomName}</label>
                            <label style="margin-top:-5px;">${item.guestRoomLevelName}</label>
                            <label style="margin-top:-5px;">￥${item.standardPriceDay}</label>
                            <input style="display: none;" value="${item.standardPrice}"/>
                            <input style="display: none;" value="${item.firstPrice}"/>
                            <input style="display: none;" value="${item.firstDuration}"/>
                            <input style="display: none;" value="${item.roomAmount}"/>
                        </button>
                    </c:if>
                    <c:if test="${item.roomStateID==7}">
                        <button onclick="suibian(this)"
                                style="width:125px;height:110px;border: 3px solid #666666; float:left;margin:2px; background:#fa0519;">
                            <input style="display: none;" value="${item.id}"/>
                            <input style="display: none;" value="${item.roomStateID}"/>
                            <label>${item.roomNumber}</label>
                            <label style="margin-top:-5px;">${item.roomName}</label>
                            <label style="margin-top:-5px;">${item.guestRoomLevelName}</label>
                            <label style="margin-top:-5px;">￥${item.standardPriceDay}</label>
                            <input style="display: none;" value="${item.standardPrice}"/>
                            <input style="display: none;" value="${item.firstPrice}"/>
                            <input style="display: none;" value="${item.firstDuration}"/>
                            <input style="display: none;" value="${item.roomAmount}"/>
                        </button>
                    </c:if>
                    <c:if test="${item.roomStateID==65}">
                        <button onclick="suibian(this)"
                                style="width:125px;height:110px;border: 3px solid #666666; float:left;margin:2px; background:rgba(255,0,242,0.95);">
                            <input style="display: none;" value="${item.id}"/>
                            <input style="display: none;" value="${item.roomStateID}"/>
                            <label>${item.roomNumber}</label>
                            <label style="margin-top:-5px;">${item.roomName}</label>
                            <label style="margin-top:-5px;">${item.guestRoomLevelName}</label>
                            <label style="margin-top:-5px;">￥${item.standardPriceDay}</label>
                            <input style="display: none;" value="${item.standardPrice}"/>
                            <input style="display: none;" value="${item.firstPrice}"/>
                            <input style="display: none;" value="${item.firstDuration}"/>
                            <input style="display: none;" value="${item.roomAmount}"/>
                        </button>
                    </c:if>
                </c:forEach>
            </div>
        </div>

    </div>

</div>
<script type="text/javascript">
    var shijian = "";
    var guestRoomLevelID = 0;

    /* function suibian(ss) {
         if (ss.children[1].value == 1) {
             var parentNodes = ss.parentNode.children;
             for (var i = 0; i < parentNodes.length; i++) {
                 parentNodes[i].style.borderColor = "#666666";
             }
             ss.style.borderColor = "#00ffff";
             document.getElementById("roomId").value = ss.children[0].value;
             document.getElementById("roomNameId").value = ss.children[2].textContent;
             document.getElementById("biaoZhunFangJiaId").value = ss.children[5].textContent;
             document.getElementById("zhongDianFangId").value = '￥' + ss.children[6].value;
             document.getElementById("shouDuanJiaGeId").value = '￥' + ss.children[7].value;
             document.getElementById("shouDuanShiChangId").value = ss.children[8].value;
             document.getElementById("countBed").value = ss.children[9].value;
         } else {
             alert("不是空房不可以查看房间的哦！");
         }


         var roomid = ss.children[0].value;  //获取这个
     }*/
    function suibian(ss) {
        var parentNodes = ss.parentNode.children;
        for (var i = 0; i < parentNodes.length; i++) {
            parentNodes[i].style.borderColor = "#666666";
        }
        ss.style.borderColor = "#00ffff";
        document.getElementById("roomId").value = ss.children[0].value;
        document.getElementById("roomNameId").value = ss.children[2].textContent;
        document.getElementById("biaoZhunFangJiaId").value = ss.children[5].textContent;
        document.getElementById("zhongDianFangId").value = '￥' + ss.children[6].value;
        document.getElementById("shouDuanJiaGeId").value = '￥' + ss.children[7].value;
        document.getElementById("shouDuanShiChangId").value = ss.children[8].value;
        document.getElementById("countBed").value = ss.children[9].value;
        var roomid = ss.children[0].value;  //获取这个
    }

    function pickedFunc() {
        shijian = $dp.cal.getNewDateStr();
    }

    function allroomfunction() {
        guestRoomLevelID = 0;
        tenfunction();
    }

    function onefunction() {
        guestRoomLevelID = 8;
        tenfunction();
    }

    function twofunction() {
        guestRoomLevelID = 9;
        tenfunction();
    }

    function threefunction() {
        guestRoomLevelID = 10;
        tenfunction();
    }

    function fourfunction() {
        guestRoomLevelID = 11;
        tenfunction();
    }

    function fivefunction() {
        guestRoomLevelID = 12;
        tenfunction();
    }

    function fiveTwofunction() {
        guestRoomLevelID = 77;
        tenfunction();
    }

    function sixfunction() {
        guestRoomLevelID = 13;
        tenfunction();
    }

    function sevenfunction() {
        guestRoomLevelID = 14;
        tenfunction();
    }

    function eightfunction() {
        guestRoomLevelID = 74;
        tenfunction();
    }

    function moviefunction() {
        guestRoomLevelID = 75;
        tenfunction();
    }

    function tenfunction() {
        $("#div1").empty();
        $.ajax({
            cache: false,                                             //是否使用缓存提交 如果为TRUE 会调用浏览器的缓存 而不会提交
            type: "POST",                                           //上面3行都是必须要的
            url: '${ctx}/StayRegister/guestRoomLevelSelectRoom.do',       //地址 type 带参数
            data: "guestRoomLevelID=" + guestRoomLevelID,                         // IDCardValue 自定义的。相当于name把值赋予给 他可以在servlet 获取
            async: false,                                          // 是否 异步 提交
            success: function (result) {                          // 不出现异常 进行立面方
                for (var key in result) {
                    var item = result[key];
                    if (item.roomStateID == 1) {
                        var btn = $("<button onclick='suibian(this)' style='width:125px;height:110px;border: 3px solid #666666; float:left;margin:2px; background:rgba(19,226,40,0.61);'>" +
                            "<input style='display: none;' value=" + item.id + " />" +
                            "<input style='display: none;' value=" + item.roomStateID + " />" +
                            "<label>" + item.roomNumber + "</label>" +
                            "<label style='margin-top:-5px;'>" + item.roomName + "</label>" +
                            "<label style='margin-top:-5px;'>" + item.guestRoomLevelName + "</label>" +
                            "<label style='margin-top:-5px;'>" + "￥" + item.standardPriceDay + "</label>" +
                            "<input style='display: none;' value=" + item.standardPrice + " />" +
                            "<input style='display: none;' value=" + item.firstPrice + " />" +
                            "<input style='display: none;' value=" + item.firstDuration + " />" +
                            "<input style='display: none;' value=" + item.roomAmount + " />" +
                            "</button>")
                        $("#div1").append(btn);
                    }
                    if (item.roomStateID == 2) {
                        var btn = $("<button onclick='suibian(this)' style='width:125px;height:110px;border: 3px solid #666666; float:left;margin:2px; background:#DDDDDD;'>" +
                            "<input style='display: none;' value=" + item.id + " />" +
                            "<input style='display: none;' value=" + item.roomStateID + " />" +
                            "<label>" + item.roomNumber + "</label>" +
                            "<label style='margin-top:-5px;'>" + item.roomName + "</label>" +
                            "<label style='margin-top:-5px;'>" + item.guestRoomLevelName + "</label>" +
                            "<label style='margin-top:-5px;'>" + "￥" + item.standardPriceDay + "</label>" +
                            "<input style='display: none;' value=" + item.standardPrice + " />" +
                            "<input style='display: none;' value=" + item.firstPrice + " />" +
                            "<input style='display: none;' value=" + item.firstDuration + " />" +
                            "<input style='display: none;' value=" + item.roomAmount + " />" +
                            "</button>")
                        $("#div1").append(btn);
                    }
                    if (item.roomStateID == 4) {
                        var btn = $("<button onclick='suibian(this)' style='width:125px;height:110px;border: 3px solid #666666; float:left;margin:2px; background:#99FFFF;'>" +
                            "<input style='display: none;' value=" + item.id + " />" +
                            "<input style='display: none;' value=" + item.roomStateID + " />" +
                            "<label>" + item.roomNumber + "</label>" +
                            "<label style='margin-top:-5px;'>" + item.roomName + "</label>" +
                            "<label style='margin-top:-5px;'>" + item.guestRoomLevelName + "</label>" +
                            "<label style='margin-top:-5px;'>" + "￥" + item.standardPriceDay + "</label>" +
                            "<input style='display: none;' value=" + item.standardPrice + " />" +
                            "<input style='display: none;' value=" + item.firstPrice + " />" +
                            "<input style='display: none;' value=" + item.firstDuration + " />" +
                            "<input style='display: none;' value=" + item.roomAmount + " />" +
                            "</button>")
                        $("#div1").append(btn);
                    }
                    if (item.roomStateID == 5) {
                        var btn = $("<button onclick='suibian(this)' style='width:125px;height:110px;border: 3px solid #666666; float:left;margin:2px; background:#BBBB00;'>" +
                            "<input style='display: none;' value=" + item.id + " />" +
                            "<input style='display: none;' value=" + item.roomStateID + " />" +
                            "<label>" + item.roomNumber + "</label>" +
                            "<label style='margin-top:-5px;'>" + item.roomName + "</label>" +
                            "<label style='margin-top:-5px;'>" + item.guestRoomLevelName + "</label>" +
                            "<label style='margin-top:-5px;'>" + "￥" + item.standardPriceDay + "</label>" +
                            "<input style='display: none;' value=" + item.standardPrice + " />" +
                            "<input style='display: none;' value=" + item.firstPrice + " />" +
                            "<input style='display: none;' value=" + item.firstDuration + " />" +
                            "<input style='display: none;' value=" + item.roomAmount + " />" +
                            "</button>")
                        $("#div1").append(btn);
                    }
                    if (item.roomStateID == 6) {
                        var btn = $("<button onclick='suibian(this)' style='width:125px;height:110px;border: 3px solid #666666; float:left;margin:2px; background:#FF7744;'>" +
                            "<input style='display: none;' value=" + item.id + " />" +
                            "<input style='display: none;' value=" + item.roomStateID + " />" +
                            "<label>" + item.roomNumber + "</label>" +
                            "<label style='margin-top:-5px;'>" + item.roomName + "</label>" +
                            "<label style='margin-top:-5px;'>" + item.guestRoomLevelName + "</label>" +
                            "<label style='margin-top:-5px;'>" + "￥" + item.standardPriceDay + "</label>" +
                            "<input style='display: none;' value=" + item.standardPrice + " />" +
                            "<input style='display: none;' value=" + item.firstPrice + " />" +
                            "<input style='display: none;' value=" + item.firstDuration + " />" +
                            "<input style='display: none;' value=" + item.roomAmount + " />" +
                            "</button>")
                        $("#div1").append(btn);
                    }
                    if (item.roomStateID == 7) {
                        var btn = $("<button onclick='suibian(this)' style='width:125px;height:110px;border: 3px solid #666666; float:left;margin:2px; background:#FF0088;'>" +
                            "<input style='display: none;' value=" + item.id + " />" +
                            "<input style='display: none;' value=" + item.roomStateID + " />" +
                            "<label>" + item.roomNumber + "</label>" +
                            "<label style='margin-top:-5px;'>" + item.roomName + "</label>" +
                            "<label style='margin-top:-5px;'>" + item.guestRoomLevelName + "</label>" +
                            "<label style='margin-top:-5px;'>" + "￥" + item.standardPriceDay + "</label>" +
                            "<input style='display: none;' value=" + item.standardPrice + " />" +
                            "<input style='display: none;' value=" + item.firstPrice + " />" +
                            "<input style='display: none;' value=" + item.firstDuration + " />" +
                            "<input style='display: none;' value=" + item.roomAmount + " />" +
                            "</button>")
                        $("#div1").append(btn);
                    }
                    if (item.roomStateID == 65) {
                        var btn = $("<button onclick='suibian(this)' style='width:125px;height:110px;border: 3px solid #666666; float:left;margin:2px; background:#FF00FF;'>" +
                            "<input style='display: none;' value=" + item.id + " />" +
                            "<input style='display: none;' value=" + item.roomStateID + " />" +
                            "<label>" + item.roomNumber + "</label>" +
                            "<label style='margin-top:-5px;'>" + item.roomName + "</label>" +
                            "<label style='margin-top:-5px;'>" + item.guestRoomLevelName + "</label>" +
                            "<label style='margin-top:-5px;'>" + "￥" + item.standardPriceDay + "</label>" +
                            "<input style='display: none;' value=" + item.standardPrice + " />" +
                            "<input style='display: none;' value=" + item.firstPrice + " />" +
                            "<input style='display: none;' value=" + item.firstDuration + " />" +
                            "<input style='display: none;' value=" + item.roomAmount + " />" +
                            "</button>")
                        $("#div1").append(btn);
                    }
                }
            },
            error: function (data) {

            }

        });
    }

    function login() {
        location.href = "${ctx}/Login/tologin.do";

    }

</script>
</body>
</html>
