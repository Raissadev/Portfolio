import React, { useState } from "react";
import api from '../services/api';
import { MailProperty, MailPattern } from "../@types/mail";
import { Col, Row, Modal, Layout, Menu, Typography, Form, Input, Alert } from 'antd';

const { Header } = Layout;
const { Title } = Typography;
const { TextArea } = Input;

function Head(): any
{
    const [isModalVisible, setIsModalVisible] = useState(false);
    const [isVisible, setIsVisible] = useState(false);

    const showModal = () => {
      setIsModalVisible(true);
    };
  
    const handleOk = async () => {
        await api.post("/mail", mail)
        .then( (response: any) => {
            setIsModalVisible(false);
            setIsVisible(true);
        })
        .catch( (err: any) => {
        })
        mail.response = "Email successfully sent";
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
                                    <a href="#home">Home</a>
                                </Menu.Item>
                                <Menu.Item key="work">
                                    <a href="#work">Work</a>
                                </Menu.Item>
                                <Menu.Item key="media">
                                    <a href="#medias">Medias</a>
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
                                placeholder="Subject"
                                type="text"
                                onChange={(val: any) => setMail({ ...mail, subject: val.target.value })}
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

                <Alert
                    message={mail.response}
                    type="success"
                    showIcon
                    closable={true}
                    style={{ display: isVisible ? 'flex' : 'none' }}
                />
            </Layout>
        </>
    );
}

export default Head;