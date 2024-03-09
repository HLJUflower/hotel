package com.zwb.controller;

import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.List;

import javax.servlet.http.HttpSession;

import com.zwb.pojo.*;
import com.zwb.service.*;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.servlet.ModelAndView;

/**
 * 用户登录
 * @Description:
 */
@Controller
@RequestMapping("/Login")
public class Login {

    @Autowired
    private UserService userService;

    @Autowired
    private StayRegisterService stayRegisterService;

    @Autowired
    private LoginLogService loginLogService;
    @Autowired
    private RoomSetService roomSetService;

    @Autowired
    private UserServieAuto userServieAuto;

    @RequestMapping("/tologin")
    public String tologin() {

        return "/login/login";
    }

    @RequestMapping("/quit")
    public ModelAndView tologin(HttpSession session) {
        Loginlog loginlog = (Loginlog) session.getAttribute("login");
        Date date = new Date();
        Loginlog log = new Loginlog();
        BeanUtils.copyProperties(loginlog, log);
        SimpleDateFormat format = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        log.setExittime(format.format(date));
//        System.out.println(log.getExittime());
        loginLogService.updateLoginLogById(log);
//        session.removeAttribute("login");
        session.invalidate();

        List<RoomSetPo> list = roomSetService.selectAll();
        ModelAndView mv = null;
        mv = new ModelAndView("redirect:/index.jsp");
        mv.addObject("list", list);
        return mv;
    }

    @RequestMapping("/tomain")
    public ModelAndView tomain(UserPo user, HttpSession session) {
        System.out.println("登录！！！");
        ModelAndView mv = null;
        double zongFeiYongOne = 0;
        double zongFeiYongTwo = 0;
        UserPo u = userService.selectLogin(user);
        List<StayRegisterPo> list = stayRegisterService.selectAll();
        for (int i = 0; i < list.size(); i++) {
            if (list.get(i).getReceiveTargetID() == 2) {
                zongFeiYongOne += list.get(i).getSumConst();
            } else {
                zongFeiYongTwo += list.get(i).getSumConst();
            }
        }

        if (u != null) {
            if ("admin".equals(u.getUserName())) {
                mv = new ModelAndView("/main/main_admin");
            } else {
                mv = new ModelAndView("/main/main");
            }
            mv.addObject("user", u);
            if (session.getAttribute("login") == null) {
                Loginlog loginlog = new Loginlog();
                loginlog.setLoginname(u.getUserName());
                loginlog.setExittime(null);
                Date date = new Date();
                SimpleDateFormat format = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
                loginlog.setLogintime(format.format(date));
                session.setAttribute("login", loginlog);
                loginLogService.saveLoginLog(loginlog);
            }
        } else {
            mv = new ModelAndView("/login/login");
            System.out.println("登录失败");
        }
        mv.addObject("zongFeiYongOne", zongFeiYongOne);
        mv.addObject("zongFeiYongTwo", zongFeiYongTwo);
        return mv;
    }


}
