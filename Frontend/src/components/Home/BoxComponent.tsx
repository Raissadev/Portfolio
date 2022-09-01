import React, { useState } from 'react';
import { Row, Col, Layout, Typography, Button, Popover, Divider } from 'antd';
import { CloudDownloadOutlined, RadarChartOutlined, HeatMapOutlined } from '@ant-design/icons';

const { Title } = Typography;

function BoxComponent(): any
{
    return(
        <>
            <Layout className="boxes-component default">
                <Row justify="space-around">
                    <Col span={6}>
                        <Popover
                            content="Linux user, I've used distros like: Linux mint, Ubuntu and currently Manjaro"
                            title="Operational system"
                            trigger="hover"
                        >
                            <Button>
                                <CloudDownloadOutlined />
                            </Button>
                        </Popover>
                        <Divider />
                    </Col>
                    <Col span={6}>
                        <Popover
                            content="I've worked with different technologies. In addition, the language that I introduced myself to the world of
                                computing was PHP, but I expanded my arsenal by adding stacks such as: JavaScript/TypeScript, Node, C, python, SQL and Golang...
                                I also like to use technologies aimed at devops like Docker .
                            "
                            title="Technologies"
                            trigger="hover"
                        >
                            <Button>
                                <RadarChartOutlined />
                            </Button>
                        </Popover>
                        <Divider />
                    </Col>
                    <Col span={6}>
                        <Popover
                            content="I find it easy to learn, in part because I'm constantly studying.
                                However, I do not have very good communication, as I am a shy person. 🙂
                            "
                            title="Soft Skills"
                            trigger="hover"
                        >
                            <Button>
                                <HeatMapOutlined />
                            </Button>
                        </Popover>
                        <Divider />
                    </Col>
                </Row>
            </Layout>
        </>
    );
}

export default BoxComponent;