package com.zwb.mapper;

import com.zwb.pojo.CommodityPo;
import org.apache.ibatis.annotations.Param;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface CommodityMapper {

    public int deleteById(Integer id);


    public int insertAll(CommodityPo commodityPo);


    public CommodityPo selectById(Integer id);


    public int updateById(CommodityPo commodityPo);


    //分页模糊查询
    public List<CommodityPo> pageFuzzyselect(@Param("commodityName") String commodityName,
                                             @Param("commodityTypeID") int commodityTypeID, @Param("start") int start, @Param("pageSize") int pageSize);
    //分页模糊查询
    public List<CommodityPo> pageFuzzySelectAll(@Param("commodityName") String commodityName,
                                             @Param("start") int start, @Param("pageSize") int pageSize);

    //分页模糊查询总条数
    public int countFuzzyselect(@Param("commodityName") String commodityName, @Param("commodityTypeID") int commodityTypeID);

    //分页模糊查询总条数
    public int countFuzzySelectAll(@Param("commodityName") String commodityName);

    //ajax 验证是否存在 此商品
    public int selectYZ(String commodityName);


    //无分页的模糊查询  非本派所用
    public List<CommodityPo> fuzzySelect(@Param("commodityName") String commodityName,
                                         @Param("commodityTypeID") int commodityTypeID);
    public List<CommodityPo> fuzzySelectAll(@Param("commodityName") String commodityName);

}
