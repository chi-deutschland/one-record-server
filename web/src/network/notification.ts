import axiosInstance from './axiosIndex';

export async function SubscribeTopic(token: string, topic: string): Promise<void> {
  await axiosInstance
    .post('/sub', {
      token,
      topic,
    })
    .then((a) => {
      console.log(a);
    })
    .catch((a) => {
      console.log(a);
    });
}
