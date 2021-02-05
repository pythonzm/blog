<template>
  <div class="dashboard-container">
    <panel-group />
    <el-row style="background:#fff;padding:16px 16px 0;margin-bottom:32px;">
      <line-chart />
    </el-row>
    <el-row :gutter="32">
      <el-col :xs="24" :sm="24" :lg="8">
        <div class="chart-wrapper">
          <CategoryPieChart />
        </div>
      </el-col>
      <el-col :xs="24" :sm="24" :lg="8">
        <div class="chart-wrapper">
          <TagPieChart />
        </div>
      </el-col>
      <el-col :xs="24" :sm="24" :lg="8">
        <div class="chart-wrapper">
          <VisitorPieChart />
        </div>
      </el-col>
      <el-row :gutter="8">
        <el-col
          :xs="{ span: 24 }"
          :sm="{ span: 24 }"
          :md="{ span: 24 }"
          :lg="{ span: 12 }"
          :xl="{ span: 12 }"
          style="padding-right:8px;margin-bottom:30px;"
        >
          <article-rank />
        </el-col>
        <el-col
          :xs="{ span: 24 }"
          :sm="{ span: 24 }"
          :md="{ span: 24 }"
          :lg="{ span: 12 }"
          :xl="{ span: 12 }"
          style="padding-right:8px;margin-bottom:30px;"
        >
          <recent-comments />
        </el-col>
      </el-row>
    </el-row>
  </div>
</template>

<script>
import PanelGroup from "./components/PanelGroup";
import LineChart from "./components/LineChart";
import CategoryPieChart from "./components/CategoryPieChart";
import TagPieChart from "./components/TagPieChart";
import VisitorPieChart from "./components/VisitorPieChart";
import ArticleRank from "./components/ArticleRank";
import RecentComments from "./components/RecentComments";

export default {
  name: "Dashboard",
  components: {
    PanelGroup,
    LineChart,
    CategoryPieChart,
    TagPieChart,
    VisitorPieChart,
    ArticleRank,
    RecentComments
  }
};
</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
