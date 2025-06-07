import styles from './DashboardPage.module.css';
import CurrentPrice from '@/components/CurrentPrice';

const DashboardPage = () => {
    return (
        <div className={styles.container}>
            <h1>Bot Dashboard</h1>
            <div className={styles.botList}>
                <div className={styles.botCard}>
                    <h2>bot-sma</h2>
                    <CurrentPrice />
                </div>
                {/* 他のbotも追加予定 */}
            </div>
        </div>
    );
};

export default DashboardPage;
