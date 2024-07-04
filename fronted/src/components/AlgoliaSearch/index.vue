<template>
  <div id="app-search">
    <div id="searchbox" />
    <div id="hits" />
    <div id="pagination" />
  </div>
</template>

<script>
import algoliasearch from 'algoliasearch/lite'
import instantsearch from 'instantsearch.js'
import { searchBox, hits, pagination } from 'instantsearch.js/es/widgets'
import 'instantsearch.css/themes/satellite.css'

const searchClient = algoliasearch('Q0SQS2P524', '5a7f4455a1571eb040b804b62ac44b00')
const search = instantsearch({
  indexName: 'blog_article',
  searchClient: searchClient,
  insights: true,
  searchFunction(helper) {
    // 如果搜索框为空，不执行搜索
    if (helper.state.query === '') {
      return
    }
    // 否则执行搜索
    helper.search()
  }
})
export default {
  name: 'AlgoliaSearch',
  mounted() {
    search.addWidgets([
      searchBox({
        container: '#searchbox',
        placeholder: '输入关键字，然后回车',
        autofocus: true,
        searchAsYouType: false,
        showReset: true,
        showSubmit: true,
        showLoadingIndicator: true,
        queryHook(query, search) {
          if (!query.trim()) {
            return false
          }
          search(query)
        }
      }),
      hits({
        container: '#hits',
        templates: {
          item(hit, { html, components }) {
            return html`
              <a href='${document.location.origin}/articles/?id=${hit.id}'>
                <h2>${components.Highlight({ hit, attribute: 'title' })}</h2>
                <p>${components.Snippet({ hit, attribute: 'content' })}...</p>
              </a>
            `
          },
          empty: ({ query }, { html }) =>
            html`<div>未找到 ${query} 相关内容</div>`
        }
      }),
      pagination({
        container: '#pagination'
      })
    ])
    search.start()
  }
}
</script>

<style>
.ais-SearchBox { margin: 1em 0; }
.ais-Hits-item:hover {
  background-color: rgb(218 215 236)
}
.ais-Pagination-list {
  justify-content: center;
}
#pagination {
  padding: 10px;
}
</style>
