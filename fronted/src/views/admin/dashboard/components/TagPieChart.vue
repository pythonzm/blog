<template>
  <div :class="className" :style="{ height: height, width: width }" />
</template>

<script>
import { fetchArticleCountByTag } from "@/api/article";
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
      fetchArticleCountByTag().then(response => {
        let data = response.data;
        let names = new Array();
        for (const item of data) {
          names.push(item.name);
        }
        this.chart = echarts.init(this.$el, "macarons");

        this.chart.setOption({
          tooltip: {
            show: true,
            formatter: function (params) {
              return params.data.name + ": " + params.data.value;
            }
          },
          series: [
            {
              name: "标签汇总",
              type: "graph",
              layout: 'force',
              roam: true,
              label: {
                show: true,
                position: 'right',
                formatter: '{b}',
                fontSize: 10
              },
              force: {
                repulsion: 120,  // 节点之间的斥力
                gravity: 0.03,
                edgeLength: [50, 150]
              },
              edgeSymbol: ['none', 'arrow'],
              edgeSymbolSize: 5,
              itemStyle: {
                borderColor: '#fff',
                borderWidth: 1,
                shadowBlur: 5,
                shadowColor: 'rgba(0, 0, 0, 0.1)'
              },
              lineStyle: {
                color: 'rgba(150, 150, 150, 0.3)',
                width: 1
              },
              data: data.map(item => ({
                ...item,
                symbolSize: 10 + item.value * 0.8,  // 节点大小跟value挂钩
                draggable: true
              })),
            }
          ]
        });
      });
    }
  }
};
</script>
