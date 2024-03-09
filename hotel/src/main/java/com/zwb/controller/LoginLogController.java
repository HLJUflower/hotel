package com.zwb.controller;

import com.zwb.page.Page;
import com.zwb.pojo.Loginlog;
import com.zwb.service.LoginLogService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.servlet.ModelAndView;

import javax.servlet.http.HttpServletRequest;


/**
 * 日志管理类  历史入住记录  个人入住历史
 * @Description:
 */
@RestController
@RequestMapping("/loginLog")
public class LoginLogController {
    @Autowired
    private LoginLogService loginLogService;

    @Value("${pwd}")
    private String pwd;

    //分页和模糊查询
    @RequestMapping("/tolist")
    public ModelAndView list(HttpServletRequest request, Integer currentPage, String txtname) {
        ModelAndView mv = null;
        mv = new ModelAndView("/loginLog/list");
        Page<Loginlog> vo = new Page<Loginlog>();
        if (currentPage == null) {
            currentPage = 1;
        } else if (currentPage == 0) {
            currentPage = 1;
        }
        if (txtname == null) {
            txtname = "";
        }
        vo.setCurrentPage(currentPage);
        vo = this.loginLogService.pageFuzzyselect(txtname, vo);
        mv.addObject("list", vo);
        mv.addObject("txtname", txtname);
        return mv;
    }

    @RequestMapping("/findLoginNameByuserName")
    public int findLoginNameByuserName(String userName, String userId) {
        Loginlog loginlog = loginLogService.selectOne(Integer.parseInt(userId));
        if (loginlog.getLoginname().equals(userName)) {
            return 1;
        }
        return 0;
    }

    @RequestMapping("/toupdate")
    public ModelAndView toupdate(String id) {
        ModelAndView mv = null;
        mv = new ModelAndView("/loginLog/update");
        Loginlog loginlog = loginLogService.selectOne(Integer.parseInt(id));
        mv.addObject("log", loginlog);
        return mv;
    }

    @RequestMapping("/update")
    public ModelAndView update(Loginlog loginlog) {
        ModelAndView mv = null;
        loginLogService.updateLoginLogById(loginlog);
        mv = new ModelAndView("redirect:/loginLog/tolist.do");
        return mv;
    }

    @RequestMapping("/isAdmin")
    public boolean isAdmin(String adminPwd) {
//        System.out.println(pwd);
        if (adminPwd.equals(pwd)) {
            return true;
        } else {
            return false;
        }
    }

    @RequestMapping("/delete")
    public ModelAndView delete(String ids) {
        ModelAndView mv = null;
        String[] FenGe = ids.split(",");
        for (int i = 0; i < FenGe.length; i++) {
            loginLogService.deleteByUserId(Integer.parseInt(FenGe[i]));
        }
        mv = new ModelAndView("redirect:/loginLog/tolist.do");
        return mv;
    }
}
