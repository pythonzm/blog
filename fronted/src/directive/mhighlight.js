import hljs from "highlight.js";
import "highlight.js/styles/monokai-sublime.css";

let Highlight = {};
Highlight.install = function(Vue, options) {
  Vue.directive("mhighlight", function(el) {
    let blocks = el.querySelectorAll("pre code");
    blocks.forEach(block => {
      hljs.highlightBlock(block);
    });
  });
};

export default Highlight;
