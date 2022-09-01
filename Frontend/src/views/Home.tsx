import { Layout } from 'antd';
import Head from "../components/Header";
import BannerComponent from '../components/Home/BannerComponent';
import BoxComponent from '../components/Home/BoxComponent';
import InterestsComponent from '../components/Home/InterestsComponent';
import MessageComponent from '../components/Home/MessageComponent';

function Home(): any
{
    return(
        <>
            <Head />

            <Layout>
                <BannerComponent />
            </Layout>

            <Layout>
                <MessageComponent />
            </Layout>

            <Layout>
                <BoxComponent />
            </Layout>

            <Layout>
                <InterestsComponent />
            </Layout>
        </>
    );
}

export default Home;