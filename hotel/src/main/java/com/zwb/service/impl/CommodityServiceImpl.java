package com.zwb.service.impl;

import java.util.List;

import com.zwb.mapper.CommodityMapper;
import com.zwb.page.Page;
import com.zwb.pojo.CommodityPo;
import com.zwb.service.CommodityService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;



@Transactional
@Service(value="commodityService")
public class CommodityServiceImpl implements CommodityService {

	@Autowired
	private CommodityMapper commodityMapper;
	
	@Override
	public int deleteById(Integer id) {
		return commodityMapper.deleteById(id);
	}

	@Override
	public int insertAll(CommodityPo commodityPo) {
		return commodityMapper.insertAll(commodityPo);
	}

	@Override
	public CommodityPo selectById(Integer id) {
		return commodityMapper.selectById(id);
	}

	@Override
	public int updateById(CommodityPo commodityPo) {
		return commodityMapper.updateById(commodityPo);
	}

	@Override
	public Page<CommodityPo> pageFuzzyselect(String commodityName,
											 int commodityTypeID, Page<CommodityPo> vo) {
		int start=0;
		if (vo.getCurrentPage()>1) {
			start=(vo.getCurrentPage()-1)*vo.getPageSize();
		}
		List<CommodityPo> list= commodityMapper.pageFuzzyselect(commodityName, commodityTypeID, start, vo.getPageSize());
		vo.setResult(list);
		int count= commodityMapper.countFuzzyselect(commodityName, commodityTypeID);
		vo.setTotal(count);
		return vo;
	}

	@Override
	public Page<CommodityPo> pageFuzzySelectAll(String commodityName, Page<CommodityPo> vo) {
		int start=0;
		if (vo.getCurrentPage()>1) {
			start=(vo.getCurrentPage()-1)*vo.getPageSize();
		}
		List<CommodityPo> list= commodityMapper.pageFuzzySelectAll(commodityName, start, vo.getPageSize());
		vo.setResult(list);
		int count= commodityMapper.countFuzzySelectAll(commodityName);
		vo.setTotal(count);
		return vo;
	}

	@Override
	public List<CommodityPo> fuzzySelect(String commodityName,int commodityTypeID) {
		if(commodityTypeID==76){
			return commodityMapper.fuzzySelectAll(commodityName);
		}
		return commodityMapper.fuzzySelect(commodityName, commodityTypeID);
	}

	@Override
	public int selectYZ(String commodityName) {
		return this.commodityMapper.selectYZ(commodityName);
	}

}
