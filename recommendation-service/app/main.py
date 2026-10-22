from fastapi import FastAPI, Query
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel

app = FastAPI(title="SEA Learning API", version="1.0.0")
app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://localhost:4173", "http://localhost:5173"],
    allow_credentials=False,
    allow_methods=["GET"],
    allow_headers=["Content-Type"],
)


class LearningPath(BaseModel):
    grade: int
    stage: str
    topics: list[str]
    next_steps: list[str]


PATHS: dict[str, dict[str, list[str]]] = {
    "primary": {
        "topics": ["Чтение и грамотная речь", "Математика и логика", "География и интерес к природе", "Английский через игры"],
        "next_steps": ["Читайте о морях и кораблях вместе со взрослыми", "Посещайте музеи и дни открытых дверей", "Развивайте командную работу и любознательность"],
    },
    "middle": {
        "topics": ["Математика", "Физика и информатика", "География", "Английский язык", "Основы первой помощи и безопасности"],
        "next_steps": ["Попробуйте профориентационные мероприятия", "Узнайте требования выбранного училища", "Поддерживайте устойчивую учебную нагрузку"],
    },
    "upper": {
        "topics": ["Профильные предметы по требованиям программы", "Английский язык", "Физика и математика для технических направлений", "Документы и сроки приёма"],
        "next_steps": ["Сверьте правила приёма на официальном сайте", "Уточните медицинские требования у комиссии", "Составьте список вопросов для приёмной комиссии"],
    },
    "student": {
        "topics": ["Профильные дисциплины", "Тренажёрная и практическая подготовка", "Морской английский", "Охрана труда и безопасность"],
        "next_steps": ["Планируйте практику с учебной частью", "Ведите журнал компетенций", "Проверяйте актуальность отраслевых требований"],
    },
}


def resolve_stage(grade: int) -> tuple[str, str]:
    if grade <= 4:
        return "Начальная школа", "primary"
    if grade <= 9:
        return "Средняя школа", "middle"
    if grade <= 11:
        return "Старшая школа", "upper"
    return "Студент или взрослый", "student"


@app.get("/health")
def health() -> dict[str, str]:
    return {"status": "ok"}


@app.get("/api/v1/learning-path", response_model=LearningPath)
def learning_path(grade: int = Query(default=8, ge=1, le=100)) -> LearningPath:
    stage, key = resolve_stage(grade)
    return LearningPath(grade=grade, stage=stage, **PATHS[key])
