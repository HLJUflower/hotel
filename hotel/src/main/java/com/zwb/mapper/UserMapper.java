package com.zwb.mapper;


import com.zwb.pojo.UserPo;
import org.springframework.stereotype.Repository;

@Repository
public interface UserMapper {
	
	public UserPo selectLogin(UserPo user);

	UserPo selectUserById(int id);

	public UserPo findIdByuserName(String userName);
	
}
