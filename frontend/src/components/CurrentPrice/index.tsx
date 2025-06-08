import { useEffect, useState } from 'react';

type Rate = {
    tag: string;
    amount: number;
};

const CurrentPrice = () => {
    const [rates, setRates] = useState<Rate[]>([]);

    useEffect(() => {
        const wsUrl = `${window.location.origin.replace(/^http/, 'ws')}/ws/price`;
        console.log('[WebSocket] CONNECTING TO:', wsUrl);

        const ws = new WebSocket(wsUrl);

        ws.onopen = () => {
            console.log('[WebSocket] OPENED');
        };

        ws.onmessage = (event) => {
            try {
                const data = JSON.parse(event.data);

                if (Array.isArray(data.rates)) {
                    const validRates = data.rates.filter(
                        (r: any) => typeof r.tag === 'string' && typeof r.amount === 'number'
                    );
                    setRates(validRates);
                } else {
                    console.warn('[WebSocket] No rates array:', data);
                }
            } catch (err) {
                console.error('[WebSocket] JSON parse error:', err);
            }
        };

        ws.onclose = (e) => {
            console.warn('[WebSocket] CLOSED', e);
        };

        return () => {
            console.log('[WebSocket] CLEANUP');
            ws.close();
        };
    }, []);

    return (
        <div style={{ padding: '1rem' }}>
            <h2 style={{ fontSize: '1.5rem', marginBottom: '1rem' }}>Exchange Rates (Base: JPY)</h2>
            {rates.length > 0 ? (
                <ul style={{ listStyle: 'none', padding: 0 }}>
                    {rates.map((rate) => (
                        <li key={rate.tag} style={{ fontSize: '1.25rem', marginBottom: '0.5rem' }}>
                            {rate.tag}: {rate.amount.toFixed(6)}
                        </li>
                    ))}
                </ul>
            ) : (
                <div>Loading exchange rates...</div>
            )}
        </div>
    );
};

export default CurrentPrice;
