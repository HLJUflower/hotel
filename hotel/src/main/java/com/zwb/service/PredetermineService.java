package com.zwb.service;

import java.sql.Timestamp;
import java.util.List;

import com.zwb.page.Page;
import com.zwb.pojo.PredeterminePo;


public interface PredetermineService {
	
	//新增预订
	public int insertAll(PredeterminePo predeterminePo);
	
	public Page<PredeterminePo> pageFuzzyselect(String receiveTeamName, String passengerName,
                                                int predetermineStateID, Page<PredeterminePo> vo);
	
	//id查询
    public PredeterminePo findById(Integer id);
    
    //根据 旅客id 来查询
    public List<PredeterminePo> findLvKeId(Integer id);
    
    //根据 团队id 来查询 
    public List<PredeterminePo> findTeamId(Integer id);
    
    //id删除
    public int deleteById(Integer id);
    
    //查询全部数据 无分页
    public List<PredeterminePo> selectAll();
    
    //修改预订到时提示
    public int updateRemind(Timestamp arriveTime,Integer id);
    
    //修改预订状态
    public int updatePredetermineStateID(Timestamp arriveTime,Integer id);

}
