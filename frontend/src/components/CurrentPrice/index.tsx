import { useEffect, useState } from 'react';

const CurrentPrice = () => {
    const [price, setPrice] = useState<number | null>(null);
    const [pair, setPair] = useState<string>('');

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
                console.log('[WebSocket] Parsed:', data);

                // たとえば: { rates: [ { pair: "USD/JPY", amount: 140.42, timestamp: 123456789 } ] }
                if (Array.isArray(data.rates) && data.rates.length > 0) {
                    const latest = data.rates[0];

                    if (typeof latest.amount === 'number') {
                        setPrice(latest.amount);
                    } else {
                        console.warn('[WebSocket] Missing or invalid amount:', latest);
                    }

                    if (typeof latest.pair === 'string') {
                        setPair(latest.pair);
                    }
                } else {
                    console.warn('[WebSocket] No rates found in data:', data);
                }
            } catch (err) {
                console.error('[WebSocket] JSON parse error:', err);
            }
        };

        ws.onclose = (e) => {
            console.warn('[WebSocket] CLOSED', e);
        };

        ws.onerror = (err) => {
            console.error('[WebSocket] ERROR:', err);
        };

        return () => {
            console.log('[WebSocket] CLEANUP, closing socket');
            ws.close();
        };
    }, []);

    return (
        <div style={{ fontSize: '2rem', fontWeight: 'bold' }}>
            {price !== null ? `${pair || 'Price'}: ${price.toFixed(6)}` : 'Loading...'}
        </div>
    );
};

export default CurrentPrice;
