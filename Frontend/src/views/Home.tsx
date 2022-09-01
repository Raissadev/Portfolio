import { Layout } from 'antd';
import BannerComponent from '../components/Home/BannerComponent';
import BoxComponent from '../components/Home/BoxComponent';
import InterestsComponent from '../components/Home/InterestsComponent';
import MessageComponent from '../components/Home/MessageComponent';

function Home(): any
{
    return(
        <>
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