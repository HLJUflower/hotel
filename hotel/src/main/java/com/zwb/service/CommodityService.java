package com.zwb.service;

import java.util.List;

import com.zwb.page.Page;
import com.zwb.pojo.CommodityPo;


public interface CommodityService {

    public int deleteById(Integer id);


    public int insertAll(CommodityPo commodityPo);


    public CommodityPo selectById(Integer id);


    public int updateById(CommodityPo commodityPo);


    //分页需要
    public Page<CommodityPo> pageFuzzyselect(String commodityName, int commodityTypeID, Page<CommodityPo> vo);

    //分页需要
    public Page<CommodityPo> pageFuzzySelectAll(String commodityName, Page<CommodityPo> vo);

    //无分页的模糊查询  非本派所用
    public List<CommodityPo> fuzzySelect(String commodityName, int commodityTypeID);

    //ajax 验证是否存在 此商品
    public int selectYZ(String commodityName);
}
