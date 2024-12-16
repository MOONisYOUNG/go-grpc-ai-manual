# 💫 go-py-grpc-ai-manual 
## <img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=Go&logoColor=white"/> <img src="https://img.shields.io/badge/Python-3776AB?style=flat-square&logo=Python&logoColor=white"/> <img src="https://img.shields.io/badge/Anaconda-44A8338?style=flat-square&logo=Anaconda&logoColor=white"/> <img src="https://img.shields.io/badge/LangChain-1C3C3C?style=flat-square&logo=LangChain&logoColor=white"/> <img src="https://github.com/user-attachments/assets/939c63cd-3dd9-43c2-9aee-f800ba4a22ae" height="20"> <img src="https://dbdb.io/media/logos/chroma_H600YUl.svg" height="20"> 
## 👉 Go/Python언어, gRPC 프레임워크, LLM, RAG 기법을 조합하여 만든 'Q&A형태의 AI 사용설명서'
### ✅ 프로젝트 소개
[SK네트웍스 Family AI캠프 2기 : 3번째 미니 프로젝트](https://github.com/SKNETWORKS-FAMILY-AICAMP/SKN02-3rd-5Team) 때 작업한 'RAG 기반 AI 사용 설명서'코드를 기반으로 추가 구현했습니다.  
Go 서버에서 질문을 남기면, Python 서버로 넘겨받은 뒤에 RAG(LangChain 라이브러리 + Chroma DB)로 연산한 값을 gRPC 기반으로 통신할 수 있도록 코드를 덧붙였습니다.

### 🔴 코드 동작 전에 해야 하는 작업
1. 'go-server' 폴더 위치에서 'go build' 명령어를 입력합니다.
2. 아나콘다 파이썬 가상환경을 설정하는 명령어 'conda create -n [가상환경 이름] python=3.10'를 입력합니다. 사용자 정의 값은 반드시 바꾸셔야 합니다.
3. 가상환경을 활성화하는 명령어 'conda activate [가상환경 이름]'을 입력합니다.
4. 'python-server' 폴더 위치에서 파이썬 라이브러리 환경을 설정하는 명령어'pip install -r requirements.txt'를 입력합니다.
5. '.env' 파일에서 발급받은 OpenAI API key 값을 입력합니다.

### 🟠 코드 동작 방법 (Linux 기준 작성)
1. 'python-sesrver' 폴더 위치에서 'python server.py' 명령어를 입력하여 python server를 활성화시킵니다.
2. 'go-server' 폴더 위치에서 './go-server' 명령어를 입력하여 go server를 활성화시킵니다.
3. go server에서 질문을 입력합니다. Q&A 기능을 종료하고 싶으면 'exit'을 입력하면 됩니다.

### 🟢 코드 실행 결과
![image](https://github.com/user-attachments/assets/2588681b-2542-4791-aa64-4180da81be44)
