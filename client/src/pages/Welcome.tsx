
import { useEffect } from 'react';
import { useInput } from '../Hotkey';
import { useNavigate } from 'react-router';

export const Welcome = () => {
    
    const navigate = useNavigate();

    useEffect(() => {
        const removeHotkeys = useInput((event: KeyboardEvent) => {
            if (event.key === "Escape") {
                event.preventDefault();
                event.stopPropagation();
                navigate('/echo');
            }
        });
        return () => {
            removeHotkeys();
        }
    }, []);

    const contributors = [
        "Sergey Anohin",
        "Andrey Mundirov",
        "Jaroslav Bespalov",
        "Richard Menedetter",
        "Tommi Koivula",
        "Rudi Timmermans",
    ];

    return (
        <>
            <div className="container">
                <img src="/static/Dog_1_2_19.png" className="welcome-img" />
                <div style={{ paddingBottom: 32 + 'px' }}>
                    <div style={{ paddingBottom: 8 + 'px' }} className="welcome-header">
                        <span>Golden point</span>
                    </div>
                    <div style={{ textAlign: 'center' }}>
                        <span>Version 1.2.19</span>
                    </div>
                </div>
                <div style={{ paddingBottom: 32 + 'px' }}>
                    <div style={{ paddingBottom: 8 + 'px' }} className="welcome-community">
                        <span>User Group Community</span>
                    </div>
                    <div style={{ textAlign: 'center' }} className="welcome-community-list">
                        <a href="https://t.me/golden_point_community" className="welcome-community-link">
                            https://t.me/golden_point_community
                        </a>
                    </div>
                </div>
                <div style={{ paddingBottom: 32 + 'px' }}>
                    <div style={{ paddingBottom: 8 + 'px' }} className="welcome-source">
                        <span>Source code and developing</span>
                    </div>
                    <a href="https://github.com/vit1251/golden" className="welcome-source-link">
                        https://github.com/vit1251/golden
                    </a>
                </div>
                <div style={{ paddingBottom: 32 + 'px' }}>
                    <div style={{ paddingBottom: 8 + 'px' }} className="welcome-contributor-header">
                        <span>Contributors</span>
                    </div>
                    <div style={{ textAlign: 'center' }} className="welcome-contributor-list">
                        <span>{contributors.join(", ")}</span>
                    </div>
                </div>

            </div>
        </>
    );

};
