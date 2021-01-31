<template>
  <el-row :gutter="40" class="panel-group">
    <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel" @click="jump('article')">
        <div class="card-panel-icon-wrapper icon-people">
          <svg-icon icon-class="list" class-name="card-panel-icon" />
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">
            文章总数
          </div>
          <count-to
            :start-val="0"
            :end-val="articleCount"
            :duration="2600"
            class="card-panel-num"
          />
        </div>
      </div>
    </el-col>
    <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel" @click="jump('category')">
        <div class="card-panel-icon-wrapper icon-message">
          <svg-icon icon-class="component" class-name="card-panel-icon" />
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">
            分类总数
          </div>
          <count-to
            :start-val="0"
            :end-val="categoryCount"
            :duration="3000"
            class="card-panel-num"
          />
        </div>
      </div>
    </el-col>
    <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel" @click="jump('tag')">
        <div class="card-panel-icon-wrapper icon-money">
          <svg-icon icon-class="tag" class-name="card-panel-icon" />
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">
            标签总数
          </div>
          <count-to
            :start-val="0"
            :end-val="tagCount"
            :duration="3200"
            class="card-panel-num"
          />
        </div>
      </div>
    </el-col>
    <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel">
        <div class="card-panel-icon-wrapper icon-shopping">
          <svg-icon icon-class="user" class-name="card-panel-icon" />
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">
            访问量
          </div>
          <count-to
            :start-val="0"
            :end-val="visitorCount"
            :duration="3600"
            class="card-panel-num"
          />
        </div>
      </div>
    </el-col>
  </el-row>
</template>

<script>
import { fetchArticleCount } from "@/api/article";
import { fetchCategoryCount } from "@/api/category";
import { fetchTagCount } from "@/api/tag";
import { fetchVisitorCount } from "@/api/visitor";
import CountTo from "vue-count-to";

export default {
  name: "Dashboard",
  components: {
    CountTo
  },
  data() {
    return {
      listLoading: true,
      articleCount: 0,
      categoryCount: 0,
      tagCount: 0,
      visitorCount: 0
    };
  },
  created() {
    this.getCount();
  },
  methods: {
    getCount() {
      this.listLoading = true;
      fetchArticleCount().then(response => {
        this.articleCount = response.data;
      });
      fetchCategoryCount().then(response => {
        this.categoryCount = response.data;
      });
      fetchTagCount().then(response => {
        this.tagCount = response.data;
      });
      fetchVisitorCount().then(response => {
        this.visitorCount = response.data;
      });
      this.listLoading = false;
    },
    jump(to) {
      if (to === "article") {
        this.$router.push({ name: "Article" });
      } else if (to === "category") {
        this.$router.push({ name: "Category" });
      } else if (to === "tag") {
        this.$router.push({ name: "Tag" });
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.panel-group {
  margin-top: 18px;

  .card-panel-col {
    margin-bottom: 32px;
  }

  .card-panel {
    height: 108px;
    cursor: pointer;
    font-size: 12px;
    position: relative;
    overflow: hidden;
    color: #666;
    background: #fff;
    box-shadow: 4px 4px 40px rgba(0, 0, 0, 0.05);
    border-color: rgba(0, 0, 0, 0.05);

    &:hover {
      .card-panel-icon-wrapper {
        color: #fff;
      }

      .icon-people {
        background: #40c9c6;
      }

      .icon-message {
        background: #36a3f7;
      }

      .icon-money {
        background: #f4516c;
      }

      .icon-shopping {
        background: #34bfa3;
      }
    }

    .icon-people {
      color: #40c9c6;
    }

    .icon-message {
      color: #36a3f7;
    }

    .icon-money {
      color: #f4516c;
    }

    .icon-shopping {
      color: #34bfa3;
    }

    .card-panel-icon-wrapper {
      float: left;
      margin: 14px 0 0 14px;
      padding: 16px;
      transition: all 0.38s ease-out;
      border-radius: 6px;
    }

    .card-panel-icon {
      float: left;
      font-size: 48px;
    }

    .card-panel-description {
      float: right;
      font-weight: bold;
      margin: 26px;
      margin-left: 0px;

      .card-panel-text {
        line-height: 18px;
        color: rgba(0, 0, 0, 0.45);
        font-size: 16px;
        margin-bottom: 12px;
      }

      .card-panel-num {
        font-size: 20px;
      }
    }
  }
}

@media (max-width: 550px) {
  .card-panel-description {
    display: none;
  }

  .card-panel-icon-wrapper {
    float: none !important;
    width: 100%;
    height: 100%;
    margin: 0 !important;

    .svg-icon {
      display: block;
      margin: 14px auto !important;
      float: none !important;
    }
  }
}
</style>
