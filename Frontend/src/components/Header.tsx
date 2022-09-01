import React, { useState } from "react";
import api from '../services/api';
import { MailProperty, MailPattern } from "../@types/mail";
import { Col, Row, Modal, Layout, Menu, Typography, Form, Input } from 'antd';

const { Header } = Layout;
const { Title } = Typography;
const { TextArea } = Input;

function Head(): any
{
    const [isModalVisible, setIsModalVisible] = useState(false);

    const showModal = () => {
      setIsModalVisible(true);
    };
  
    const handleOk = async () => {
        await api.post("/mail", mail)
        .then( (response: any) => {
            setIsModalVisible(false);
        })
        .catch( (err: any) => {
            mail.message = "Invalid!";
        })
    };
  
    const handleCancel = () => {
      setIsModalVisible(false);
    };

    const [scheme]: any = Form.useForm();
    const [mail, setMail]: any = useState<MailProperty>(MailPattern);

    return(
        <>
            <Layout>
                <Header>
                    <Row align="middle" justify="space-between">
                        <Col>
                            <Title level={4}>Raissadev</Title>
                        </Col>
                        <Col span={6} offset={6}>
                            <Menu mode="horizontal">
                                <Menu.Item key="home">
                                    Home
                                </Menu.Item>
                                <Menu.Item key="work">
                                    Work
                                </Menu.Item>
                                <Menu.Item key="media">
                                    Medias
                                </Menu.Item>
                                <Menu.Item key="email" onClick={showModal}>
                                    Email
                                </Menu.Item>
                            </Menu>
                        </Col>
                    </Row>
                </Header>

                <Modal title="Send Mail" visible={isModalVisible} onOk={handleOk} onCancel={handleCancel}>
                    <Form
                        name="basic"
                        labelCol={{ span: 8 }}
                        wrapperCol={{ span: 16 }}
                        initialValues={{ remember: true }}
                        onFinish={handleOk}
                        form={scheme}
                        autoComplete="off"
                    >
                        <Form.Item
                            name="name"
                            rules={[{ required: true, message: 'Please input your name!' }]}
                        >
                            <Input
                                placeholder="Name"
                                type="text"
                                onChange={(val: any) => setMail({ ...mail, name: val.target.value })}
                            />
                        </Form.Item>
                        <Form.Item
                            name="email"
                            rules={[{ required: true, message: 'Please input your email!' }]}
                        >
                            <Input
                                placeholder="Email"
                                type="email"
                                onChange={(val: any) => setMail({ ...mail, email: val.target.value })}
                            />
                        </Form.Item>
                        <Form.Item
                            name="message"
                            rules={[{ required: true, message: 'Please input your message!' }]}
                        >
                            <TextArea
                                placeholder="Message"
                                onChange={(val: any) => setMail({ ...mail, message: val.target.value })}
                                rows={4}
                            />
                        </Form.Item>
                    </Form>
                </Modal>
            </Layout>
        </>
    );
}

export default Head;