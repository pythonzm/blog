<template>
  <div :class="className" :style="{ height: height, width: width }" />
</template>

<script>
import { fetchCountByUa } from "@/api/visitor";
import echarts from "echarts";
require("echarts/theme/macarons"); // echarts theme
import resize from "./mixins/resize";

export default {
  mixins: [resize],
  props: {
    className: {
      type: String,
      default: "chart"
    },
    width: {
      type: String,
      default: "100%"
    },
    height: {
      type: String,
      default: "300px"
    }
  },
  data() {
    return {
      chart: null
    };
  },
  mounted() {
    this.$nextTick(() => {
      this.initChart();
    });
  },
  beforeDestroy() {
    if (!this.chart) {
      return;
    }
    this.chart.dispose();
    this.chart = null;
  },
  methods: {
    initChart() {
      fetchCountByUa().then(response => {
        let data = response.data;
        let names = new Array();
        for (const item of data) {
          names.push(item.name);
        }
        this.chart = echarts.init(this.$el, "macarons");

        this.chart.setOption({
          tooltip: {
            trigger: "item",
            formatter: "{a} <br/>{b} : {c} ({d}%)"
          },
          legend: {
            left: "center",
            bottom: "10",
            data: names
          },
          series: [
            {
              name: "访客汇总",
              type: "pie",
              roseType: "radius",
              radius: [15, 95],
              center: ["50%", "38%"],
              data: data,
              animationEasing: "cubicInOut",
              animationDuration: 2600
            }
          ]
        });
      });
    }
  }
};
</script>
