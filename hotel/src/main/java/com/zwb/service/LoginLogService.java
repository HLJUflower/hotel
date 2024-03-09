package com.zwb.service;

import com.zwb.page.Page;
import com.zwb.pojo.Loginlog;
import java.util.List;

/**
 * @Description:
 */
public interface LoginLogService {
    public int saveLoginLog(Loginlog loginlog);

    public List<Loginlog> selectAll();

    //分页加模糊查询
    public Page<Loginlog> pageFuzzyselect(String loginName, Page<Loginlog> vo);

    public void updateLoginLogById(Loginlog loginlog);

    public Loginlog selectOne(Integer id);

    void deleteByUserId(int parseInt);
}
