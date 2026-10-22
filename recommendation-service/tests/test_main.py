from fastapi.testclient import TestClient
from app.main import app

client = TestClient(app)


def test_health() -> None:
    response = client.get("/health")
    assert response.status_code == 200
    assert response.json() == {"status": "ok"}


def test_learning_path_for_middle_school() -> None:
    response = client.get("/api/v1/learning-path", params={"grade": 8})
    assert response.status_code == 200
    payload = response.json()
    assert payload["stage"] == "Средняя школа"
    assert "Математика" in payload["topics"]
    assert payload["next_steps"]


def test_invalid_grade_is_rejected() -> None:
    assert client.get("/api/v1/learning-path", params={"grade": 0}).status_code == 422
