module.exports = {
  title: 'POOROPS',

  /**
   * @type {boolean} true | false
   * @description Whether fix the header
   */
  fixedHeader: true,

  /**
   * @type {boolean} true | false
   * @description Whether show the logo in sidebar
   */
  sidebarLogo: true,

  /**
   * @type {boolean} true | false
   * @description Whether use the algolia search service
   */
  algoliaSearch: true,

  /**
   * @type {string}
   * @description the algolia search appID
   */
  algoliaAppID: process.env.VUE_APP_ALGOLIA_APP_ID || '',

  /**
   * @type {string}
   * @description the algolia search api key
   */
  algoliaApiKey: process.env.VUE_APP_ALGOLIA_API_KEY || '',

  /**
   * @type {string}
   * @description the algolia search index name
   */
  algoliaIndexName: process.env.VUE_APP_ALGOLIA_INDEX_NAME || 'blog_article',

  /**
   * @type {boolean} true | false
   * @description Whether to enable console anti-debugging script in production to block F12 inspector
   */
  antiDebug: true
}
