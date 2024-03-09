package com.zwb.service;


import com.zwb.pojo.UserPo;

public interface UserService {

	public UserPo selectLogin(UserPo user);
	public UserPo findIdByuserName(String userName);
	
}
