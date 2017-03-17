import Preact from 'preact'
import { renderToString } from 'react-dom/server'
import NextDocumentConstructor from 'next/constructors/document'
import flush from 'styled-jsx/server'

export class Document extends Preact.Component {
  static getInitialProps ({ pageComponent, getHead }) {
    // Likewise, you could use your own renderer here.
    const { html } = renderToString(pageComponent)
    const head = getHead()
    const styles = flush()
    return { html, styles }
  }

  return () {
    return NextDocumentConstructor(Preact, this.props)
  }
}