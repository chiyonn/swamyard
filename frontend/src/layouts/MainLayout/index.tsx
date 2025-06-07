import { Outlet } from 'react-router-dom';
import styles from './MainLayout.module.css';

const MainLayout = () => {
    return (
        <div className={styles.container}>
            <header className={styles.header}></header>
            <div className={styles.contentWrapper}>
                <aside className={styles.sidebarLeft}></aside>
                <main className={styles.main}>
                    <Outlet />
                </main>
                <aside className={styles.sidebarRight}></aside>
            </div>
            <footer className={styles.footer}></footer>
        </div>
    );
};

export default MainLayout;

