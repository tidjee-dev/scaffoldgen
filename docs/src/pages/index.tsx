import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import type { ReactNode } from 'react';
import Heading from '@theme/Heading';
import Link from '@docusaurus/Link';
import Layout from '@theme/Layout';
import clsx from 'clsx';

import styles from './index.module.css';

function HomepageHeader() {
  const {siteConfig} = useDocusaurusContext();
  return (
    <header className={clsx('hero hero--primary', styles.heroBanner)}>
      <div className="container">
        <Heading as="h1" className="hero__title">
          {siteConfig.title}
        </Heading>
        <p className="hero__subtitle">{siteConfig.tagline}</p>
        <div className={styles.buttons}>
          <Link
            className={clsx('button button--lg', styles.primary)}
            to="/docs/getting-started/installation">
            Get Started
          </Link>
          <Link
            className={clsx('button button--lg', styles.link)}
            to="https://github.com/tidjee-dev/scaffoldgen">
            View on GitHub
          </Link>
        </div>
      </div>
    </header>
  );
}

export default function Home(): ReactNode {
  const {siteConfig} = useDocusaurusContext();
  return (
    <Layout
      title={siteConfig.title}
      description="Safe, reviewable scaffold generation from structured formats">
      <HomepageHeader />
      <main>
        <section className="margin-vert--lg padding-vert--lg">
          <div className="container">
            <div className={styles.features}>
              <div>
                <h3>📋 Multiple Input Formats</h3>
                <p>Write structures in Markdown (fastest), YAML (readable), or JSON (programmatic)</p>
              </div>
              <div>
                <h3>🔒 Safety First</h3>
                <p>Generate shell scripts that you review and execute manually</p>
              </div>
              <div>
                <h3>🚀 Cross-Platform</h3>
                <p>Target Linux/macOS with Bash or Windows with PowerShell</p>
              </div>
              <div>
                <h3>🧠 Language-Aware</h3>
                <p>Automatic boilerplate for 20+ languages (Go, Python, TypeScript, Rust, Java, etc.)</p>
              </div>
              <div>
                <h3>👁️ Preview System</h3>
                <p>Visualize exactly what will be created before generating</p>
              </div>
              <div>
                <h3>🔄 Reverse Mode</h3>
                <p>Scan existing directories and generate structure files</p>
              </div>
            </div>
          </div>
        </section>
      </main>
    </Layout>
  );
}
