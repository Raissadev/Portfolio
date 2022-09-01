import { Row, Col, Layout, Typography, Divider, Image } from 'antd';

const { Title } = Typography;

function InterestsComponent(): any
{
    return(
        <>
            <Layout className="interests-component default">
                <Row>
                    <Col span={24}>
                        <Divider orientation="left" orientationMargin="0">
                            <Title level={2}>
                                DevOps
                            </Title>
                        </Divider>
                    </Col>
                    <Col>
                        <Image width={1000} src="https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png"/>
                    </Col>
                </Row>
                <Row>
                    <Col span={24}>
                        <Divider orientation="right" orientationMargin="0">
                            <Title level={2}>
                                Backend
                            </Title>
                        </Divider>
                    </Col>
                    <Col>
                        <Image width={1000} src="https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png"/>
                    </Col>
                </Row>
                <Row>
                    <Col span={24}>
                        <Divider orientation="left" orientationMargin="0">
                            <Title level={2}>
                                Frontend
                            </Title>
                        </Divider>
                    </Col>
                    <Col>
                        <Image width={1000} src="https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png"/>
                    </Col>
                </Row>
            </Layout>
        </>
    );
}

export default InterestsComponent;