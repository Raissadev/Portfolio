import React, { useState } from "react";
import { Col, Row, MenuProps, Layout, Menu, Typography } from 'antd';

const { Header } = Layout;
const { Title } = Typography;

function Head(): any
{
    const items: MenuProps['items'] = [
        { label: 'Home', key: 'home' },
        { label: 'Work', key: 'work' },
        { label: 'Contact', key: 'contact' },
        { label: 'Email', key: 'email' },
    ];

    const [current, setCurrent] = useState('mail');
      
    const onClick: MenuProps['onClick'] = e => {
        console.log('click ', e);
        setCurrent(e.key);
    };

    return(
        <>
            <Layout>
                <Header>
                    <Row align="middle" justify="space-between">
                        <Col>
                            <Title level={4}>Raissadev</Title>
                        </Col>
                        <Col span={6} offset={6}>
                            <Menu onClick={onClick} selectedKeys={[current]} mode="horizontal" items={items} />
                        </Col>
                    </Row>
                </Header>
            </Layout>
        </>
    );
}

export default Head;