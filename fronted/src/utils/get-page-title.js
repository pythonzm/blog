import defaultSettings from '@/settings'

const title = defaultSettings.title || 'POOROPS'

export default function getPageTitle(pageTitle) {
  if (pageTitle) {
    return `${pageTitle} - ${title}`
  }
  return `${title}`
}
