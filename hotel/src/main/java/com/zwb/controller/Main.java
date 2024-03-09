package com.zwb.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
/**
 * 跳板页
 * @Description:
 */
@Controller
@RequestMapping("/Main")
public class Main {
	
	@RequestMapping("/todata")
	public String tologin(){
		return "redirect:/DateView/tolist.do?type="+1;
	}

}
