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
    </el-row>
  </div>
</template>

<script>
import PanelGroup from "./components/PanelGroup";
import LineChart from "./components/LineChart";
import CategoryPieChart from "./components/CategoryPieChart";
import TagPieChart from "./components/TagPieChart";
import VisitorPieChart from "./components/VisitorPieChart";

export default {
  name: "Dashboard",
  components: {
    PanelGroup,
    LineChart,
    CategoryPieChart,
    TagPieChart,
    VisitorPieChart
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
